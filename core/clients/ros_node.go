package clients

import (
	"fmt"

	"github.com/fetchrobotics/rosgo/ros"
	"github.com/spf13/viper"
)

// NewROSNode constructor based on given config
func NewROSNode(config *viper.Viper) (ros.Node, error) {
	// Acquire configurations
	nodeName := config.GetString("ros.nodeName")
	nodeArgs := config.GetStringSlice("ros.args")

	// Create new ROS Node
	newNode, err := ros.NewNode(nodeName, nodeArgs)
	if err != nil {
		return nil, fmt.Errorf("Create ROS Node: %s", err)
	}
	newNode.Logger().SetSeverity(ros.LogLevelDebug)

	return newNode, nil
}
