package clients

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/street-bot/robot/libs/websocket"
)

// NewWSClient constructor
func NewWSClient(config *viper.Viper) (*websocket.Socket, error) {
	signalEndpoint := config.GetString("signalEndpoint")
	if signalEndpoint == "" {
		return nil, fmt.Errorf("signalEndpoint is not specified in configuration file")
	}

	client := websocket.New(signalEndpoint)
	err := client.Connect()
	return client, err
}
