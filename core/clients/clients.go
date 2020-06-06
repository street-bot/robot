package clients

import (
	"github.com/fetchrobotics/rosgo/ros"
	"github.com/street-bot/robot/libs/websocket"
)

// Clients interface to get the specific client
type Clients interface {
	WebSocket() *websocket.Socket

	StartROSNode(crashed chan bool)

	ROSPub(topic string) ros.Publisher
	AddROSPub(topic string, messageType ros.MessageType)
	RemoveROSPub(topic string) error

	ROSSub(topic string) ros.Subscriber
	AddROSSub(topic string, messageType ros.MessageType, callback interface{})
	RemoveROSSub(topic string) error
}
