package signaling

import (
	"github.com/go-resty/resty/v2"
	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	"github.com/street-bot/robot/core/realtime"
	rlog "github.com/street-bot/robot/libs/log"
)

// Signaler functions to create the WebRTC connection
type Signaler interface {
	RegisterPeerConnection(pc realtime.Connection) // This PeerConnection object should be created elsewhere
}

// RobotSignaler implements the Signaler interface to establish the WebRTC Connection
type RobotSignaler struct {
	clients    clients.Clients
	logger     rlog.Logger
	conn       realtime.Connection
	config     *viper.Viper
	http       *resty.Client
	iceServers []webrtc.ICEServer

	onOfferCb    func(string)
	onDisconnect func(rlog.Logger, *viper.Viper) error
}

// NewRobotSignaler constructor for the WebRTC signaler
func NewRobotSignaler(clients clients.Clients, logger rlog.Logger, conn realtime.Connection, config *viper.Viper) (*RobotSignaler, error) {
	newSignaler := new(RobotSignaler)
	newSignaler.clients = clients
	newSignaler.logger = logger
	newSignaler.conn = conn
	newSignaler.config = config
	newSignaler.http = resty.New()

	newSignaler.GetICEServers()
	newSignaler.registerSocketTransportCallbacks()

	return newSignaler, nil
}
