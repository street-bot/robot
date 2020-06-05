package clients

import (
	"github.com/fetchrobotics/rosgo/ros"
	socketio_client "github.com/frankgu968/go-socket.io-client"
)

// Clients interface to get the specific client
type Clients interface {
	SocketIO() *socketio_client.Client

	StartROSNode(crashed chan bool)

	ROSPub(topic string) ros.Publisher
	AddROSPub(topic string, messageType ros.MessageType)
	RemoveROSPub(topic string) error

	ROSSub(topic string) ros.Subscriber
	AddROSSub(topic string, messageType ros.MessageType, callback interface{})
	RemoveROSSub(topic string) error
}
