package signaling

import (
	socketio_client "github.com/frankgu968/go-socket.io-client"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/spf13/viper"
)

// RegisterSocketIOCallbacks will create all the handlers for SocketIO events
func registerSocketIOCallbacks(client *socketio_client.Client, logger rlog.Logger, config *viper.Viper) {
	client.On("error", OnError(logger))

	// client.On("connection", OnConnect(logger, config, client))
	client.On("connection", OnConnect(logger, config, client))

	client.On("disconnection", OnDisconnect(logger))
}

// OnError event handler constructor
func OnError(logger rlog.Logger) func(string) {
	return func(msg string) {
		logger.Warnf("SocketIO error: %s", msg)
	}
}

// OnConnect event handler constructor
func OnConnect(logger rlog.Logger, config *viper.Viper, client *socketio_client.Client) func() {
	return func() {
		logger.Infof("Connected to signaling server")
		client.Emit("/robot", config.GetString("id"))
	}
}

// OnDisconnect event handler constructor
func OnDisconnect(logger rlog.Logger) func() {
	return func() {
		logger.Errorf("Disconnected from signaling server")
	}
}
