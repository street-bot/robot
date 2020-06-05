package clients

import (
	"fmt"

	socketio_client "github.com/frankgu968/go-socket.io-client"
	"github.com/spf13/viper"
)

// NewSocketIOClient constructor based on given config
func NewSocketIOClient(config *viper.Viper) (*socketio_client.Client, error) {
	signalEndpoint := config.GetString("signalEndpoint")
	if signalEndpoint == "" {
		return nil, fmt.Errorf("signalEndpoint is not specified in configuration file")
	}

	opts := &socketio_client.Options{
		Transport: "websocket",
		WSScheme:  "ws",
	}
	newClient, err := socketio_client.NewClient(signalEndpoint, opts)
	if err != nil {
		return nil, fmt.Errorf("Create SocketIO client: %s", err.Error())
	}

	return newClient, nil
}
