package websocket

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Socket abstraction on top of WebSocket
type Socket struct {
	Conn              *websocket.Conn
	WebsocketDialer   *websocket.Dialer
	URL               string
	ConnectionOptions ConnectionOptions
	RequestHeader     http.Header
	OnBinaryMessage   func(data []byte, socket *Socket)
	OnPingReceived    func(data string, socket *Socket)
	OnPongReceived    func(data string, socket *Socket)
	OnConnected       func(socket *Socket)
	OnMessage         func(message string, socket *Socket)
	OnError           func(err error, socket *Socket)
	OnDisconnected    func(err error, socket *Socket)
	IsConnected       bool
	sendMu            *sync.Mutex // Prevent "concurrent write to websocket connection"
	receiveMu         *sync.Mutex
	pingInterval      time.Duration
}

// ConnectionOptions define properties of the websocket
type ConnectionOptions struct {
	UseCompression bool
	UseSSL         bool
	Proxy          func(*http.Request) (*url.URL, error)
	Subprotocols   []string
}

// New websocket object
func New(url string) *Socket {
	return &Socket{
		URL:           url,
		RequestHeader: http.Header{},
		ConnectionOptions: ConnectionOptions{
			UseCompression: false,
			UseSSL:         true,
		},
		WebsocketDialer: &websocket.Dialer{},
		sendMu:          &sync.Mutex{},
		receiveMu:       &sync.Mutex{},
	}
}

func (socket *Socket) setConnectionOptions() {
	socket.WebsocketDialer.EnableCompression = socket.ConnectionOptions.UseCompression
	socket.WebsocketDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: socket.ConnectionOptions.UseSSL}
	socket.WebsocketDialer.Proxy = socket.ConnectionOptions.Proxy
	socket.WebsocketDialer.Subprotocols = socket.ConnectionOptions.Subprotocols
	socket.pingInterval = 2 * time.Second
}

// Connect the websocket
func (socket *Socket) Connect() error {
	var err error
	socket.setConnectionOptions()

	socket.Conn, _, err = socket.WebsocketDialer.Dial(socket.URL, socket.RequestHeader)

	if err != nil {
		socket.IsConnected = false
		if socket.OnError != nil {
			socket.OnError(err, socket)
		}
		return err
	}

	socket.IsConnected = true
	if socket.OnConnected != nil {
		socket.OnConnected(socket)
	}

	defaultPingHandler := socket.Conn.PingHandler()
	socket.Conn.SetPingHandler(func(appData string) error {
		if socket.OnPingReceived != nil {
			socket.OnPingReceived(appData, socket)
		}
		return defaultPingHandler(appData)
	})

	defaultPongHandler := socket.Conn.PongHandler()
	socket.Conn.SetPongHandler(func(appData string) error {
		if socket.OnPongReceived != nil {
			socket.OnPongReceived(appData, socket)
		}
		return defaultPongHandler(appData)
	})

	defaultCloseHandler := socket.Conn.CloseHandler()
	socket.Conn.SetCloseHandler(func(code int, text string) error {
		result := defaultCloseHandler(code, text)
		socket.IsConnected = false
		if socket.OnDisconnected != nil {
			socket.OnDisconnected(errors.New(text), socket)
		}
		return result
	})

	// Heartbeat sender
	stop := make(chan bool) // Quit channel for heartbeat intervals
	go socket.heartbeat(stop)
	go socket.fetchLoop(stop)

	return nil
}

// WebSocket heartbeat loop after connection
func (socket *Socket) heartbeat(quit chan bool) {
	ticker := time.NewTicker(socket.pingInterval)
	for {
		select {
		case <-ticker.C:
			if err := socket.SendRaw(websocket.PingMessage, []byte{}); err != nil {
				return
			}

		// Stop the ping health check
		case <-quit:
			close(quit)
			return
		}
	}
}

// Main websocket loop to fetch messages and make calls
func (socket *Socket) fetchLoop(hbstop chan bool) {
	pongWait := socket.pingInterval + 1*time.Second // 1s of timeout leeway
	socket.Conn.SetReadDeadline(time.Now().Add(pongWait))
	socket.Conn.SetPongHandler(func(string) error {
		socket.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		socket.receiveMu.Lock()
		messageType, message, err := socket.Conn.ReadMessage()
		socket.receiveMu.Unlock()
		if err != nil {
			hbstop <- true
			socket.IsConnected = false
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
				// Not abnormal closure (server disconnected) -> must be an error!
				if socket.OnError != nil {
					socket.OnError(err, socket)
				}
			} else {
				if socket.OnDisconnected != nil {
					socket.OnDisconnected(err, socket)
				}
			}
			return
		}

		switch messageType {
		case websocket.TextMessage:
			if socket.OnMessage != nil {
				socket.OnMessage(string(message), socket)
			}
		case websocket.BinaryMessage:
			if socket.OnBinaryMessage != nil {
				socket.OnBinaryMessage(message, socket)
			}
		}
	}
}

// Send a string message
func (socket *Socket) Send(message string) {
	err := socket.SendRaw(websocket.TextMessage, []byte(message))
	if err != nil {
		return
	}
}

// SendBinary payload
func (socket *Socket) SendBinary(data []byte) {
	err := socket.SendRaw(websocket.BinaryMessage, data)
	if err != nil {
		return
	}
}

// SendRaw message type
func (socket *Socket) SendRaw(messageType int, data []byte) error {
	socket.sendMu.Lock()
	err := socket.Conn.WriteMessage(messageType, data)
	socket.sendMu.Unlock()
	return err
}

// Close the connection
func (socket *Socket) Close() {
	if socket.IsConnected {
		// Only make the graceful shutdown if the socket is still open
		err := socket.SendRaw(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		socket.Conn.Close()
		socket.IsConnected = false
		if socket.OnDisconnected != nil {
			socket.OnDisconnected(err, socket)
		}
	}
}
