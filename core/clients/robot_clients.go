package clients

import (
	"fmt"
	"time"

	"github.com/fetchrobotics/rosgo/ros"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/libs/websocket"
)

// RobotClients concrete implementation of the Clients interface
type RobotClients struct {
	ws              *websocket.Socket
	rosNode         ros.Node
	rosSpinInterval time.Duration

	rosPubs map[string]ros.Publisher
	rosSubs map[string]ros.Subscriber
}

// NewRobotClients constructor based on given config
func NewRobotClients(config *viper.Viper) (Clients, error) {
	newClientSet := new(RobotClients)
	newClientSet.rosPubs = make(map[string]ros.Publisher)
	newClientSet.rosSubs = make(map[string]ros.Subscriber)

	// Parse the node spin interval
	nodeSpinIntervalStr := config.GetString("ros.spinInterval")
	spinInterval, err := time.ParseDuration(nodeSpinIntervalStr)
	if err != nil {
		return nil, err
	}
	newClientSet.rosSpinInterval = spinInterval

	// Instantiate SocketTransport Client
	transport, err := NewWSClient(config)
	if err != nil {
		return nil, err
	}
	newClientSet.ws = transport

	// Instantiate ROS Node
	rosNode, err := NewROSNode(config)
	if err != nil {
		return nil, err
	}
	newClientSet.rosNode = rosNode

	return newClientSet, nil
}

// WebSocket accessor
func (c *RobotClients) WebSocket() *websocket.Socket {
	return c.ws
}

// StartROSNode will run ROS node loop
func (c *RobotClients) StartROSNode(crashed chan bool) {
	for c.rosNode.OK() {
		c.rosNode.SpinOnce()
		// time.Sleep(c.rosSpinInterval)
	}

	// Should not reach here during normal operations
	// Notify crashed channel
	crashed <- true
}

// ROSPub accessor
func (c *RobotClients) ROSPub(topic string) ros.Publisher {
	return c.rosPubs[topic]
}

// AddROSPub creates a new ROS publisher to the RobotClients struct
func (c *RobotClients) AddROSPub(topic string, msgType ros.MessageType) {
	c.rosPubs[topic] = c.rosNode.NewPublisher(topic, msgType)
}

// RemoveROSPub removes a publisher from the publisher map
func (c *RobotClients) RemoveROSPub(topic string) error {
	_, ok := c.rosPubs[topic]
	if ok {
		delete(c.rosPubs, topic)
	} else {
		return fmt.Errorf("Attempted to delete non-existent ROS publisher: %s", topic)
	}

	return nil
}

// ROSSub accessor
func (c *RobotClients) ROSSub(topic string) ros.Subscriber {
	return c.rosSubs[topic]
}

// AddROSSub creates a new ROS subscriber to the RobotClients struct
func (c *RobotClients) AddROSSub(topic string, msgType ros.MessageType, callback interface{}) {
	c.rosSubs[topic] = c.rosNode.NewSubscriber(topic, msgType, callback)
}

// RemoveROSSub removes a subscriber from the subscriber map
func (c *RobotClients) RemoveROSSub(topic string) error {
	_, ok := c.rosSubs[topic]
	if ok {
		delete(c.rosSubs, topic)
	} else {
		return fmt.Errorf("Attempted to delete non-existent ROS subscriber: %s", topic)
	}

	return nil
}
