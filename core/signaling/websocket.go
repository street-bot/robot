package signaling

import (
	"encoding/json"

	"github.com/street-bot/robot/core/realtime"
	"github.com/street-bot/robot/libs/websocket"
)

// registerSocketTransportCallbacks will create all the handlers for SocketIO events
func (rs *RobotSignaler) registerSocketTransportCallbacks() {
	client := rs.clients.WebSocket()
	client.OnError = rs.makeErrorHandler()

	client.OnConnected = rs.makeConnectHandler()

	client.OnDisconnected = rs.makeDisconnectHandler()

	client.OnMessage = rs.makeMessageHandler()
}

func (rs *RobotSignaler) makeErrorHandler() func(error, *websocket.Socket) {
	return func(err error, socket *websocket.Socket) {
		rs.logger.Warnf("Websocket error: %s", err.Error())
	}
}

func (rs *RobotSignaler) makeConnectHandler() func(socket *websocket.Socket) {
	return func(socket *websocket.Socket) {
		rs.logger.Infof("Connected to signaling server")
		rregMsg := NewRobotRegistrationMessage(rs.config.GetString("id"))
		msgStr, err := rregMsg.ToString()
		if err != nil {
			rs.logger.Errorf("Send robot registration: %s", err.Error())
			return
		}
		socket.Send(msgStr) // Register robot with signaler
	}
}

func (rs *RobotSignaler) makeDisconnectHandler() func(error, *websocket.Socket) {
	return func(err error, socket *websocket.Socket) {
		rs.logger.Errorf("Disconnected from signaling server")
		if err != nil {
			rs.logger.Errorf(err.Error())
		}
	}
}

func (rs *RobotSignaler) makeMessageHandler() func(string, *websocket.Socket) {
	return func(msg string, socket *websocket.Socket) {
		rs.logger.Debugf("Received message %s", msg)
		var parsedMsg BaseMessage
		if err := json.Unmarshal([]byte(msg), &parsedMsg); err != nil {
			rs.logger.Errorf("Message parse: %s", err.Error())
			return
		}
		switch parsedMsg.Type {
		case OfferType:
			// Ensure the onOfferCb function is present
			var offerMsg OfferMessage
			if err := json.Unmarshal([]byte(msg), &offerMsg); err != nil {
				rs.logger.Errorf("Message parse: %s", err.Error())
				return
			}
			rs.onOfferCb(offerMsg.Payload.SDPStr)
		}
		// TODO: Message parser and dispatcher here
	}
}

// RegisterPeerConnection listens for offers and establish connection
func (rs *RobotSignaler) RegisterPeerConnection(rtc realtime.Connection) {
	rs.onOfferCb = rs.onOffer(rtc)
	// rs.clients.SocketTransport().On("/offer", rs.onOffer(rtc))
}
