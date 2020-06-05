package signaling

import (
	"time"

	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/realtime"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/street-bot/robot/libs/websocket"
)

// registerSocketTransportCallbacks will create all the handlers for SocketIO events
func registerSocketTransportCallbacks(client *websocket.Socket, logger rlog.Logger, config *viper.Viper) {
	client.OnError = OnError(logger)

	client.OnConnected = OnConnect(logger, config)

	client.OnDisconnected = OnDisconnect(logger)
}

// OnError event handler constructor
func OnError(logger rlog.Logger) func(error, *websocket.Socket) {
	return func(err error, socket *websocket.Socket) {
		logger.Warnf("Websocket error: %s", err.Error)
	}
}

// OnConnect event handler constructor
func OnConnect(logger rlog.Logger, config *viper.Viper) func(socket *websocket.Socket) {
	return func(socket *websocket.Socket) {
		logger.Infof("Connected to signaling server")
		rregMsg := NewRobotRegistrationMessage(config.GetString("id"))
		msgStr, err := rregMsg.ToString()
		if err != nil {
			logger.Errorf("Send robot registration: %s", err.Error())
			return
		}
		socket.Send(msgStr) // Register robot with signaler
	}
}

// OnDisconnect event handler constructor
func OnDisconnect(logger rlog.Logger) func(error, *websocket.Socket) {
	return func(err error, socket *websocket.Socket) {
		logger.Errorf("Disconnected from signaling server")
		if err != nil {
			logger.Errorf(err.Error())
		}
		reconnectFunc := func() {
			if err := socket.Connect(); err != nil {
				logger.Errorf("Reconnect error: %s", err.Error())
			}
		}
		time.AfterFunc(1*time.Second, reconnectFunc) // Attempt reconnection if disconnected
	}
}

// RegisterPeerConnection listens for offers and establish connection
func (rs *RobotSignaler) RegisterPeerConnection(rtc realtime.Connection) {
	// rs.clients.SocketTransport().On("/offer", rs.onOffer(rtc))
}
