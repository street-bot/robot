package core

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	"github.com/street-bot/robot/core/realtime"
	"github.com/street-bot/robot/core/signaling"
	"github.com/street-bot/robot/libs/config"
	rlog "github.com/street-bot/robot/libs/log"
)

// Robot interface
type Robot interface {
	Start()
}

// RobotCore contains all the clients and robot-specific data
type RobotCore struct {
	logger   rlog.Logger
	config   *viper.Viper
	clients  clients.Clients
	signaler signaling.Signaler
	rtc      realtime.Connection
}

// NewRobotCore constructor
func NewRobotCore() Robot {
	newRobotCore := new(RobotCore)

	// Read configuration
	parsedConfig, err := config.Init()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	newRobotCore.config = parsedConfig

	// Set up logger
	logLevel := newRobotCore.config.GetString("log.level")
	newRobotCore.logger = rlog.NewZeroLogger(logLevel, os.Stdout)

	// Set up robot clients
	newClients, err := clients.NewRobotClients(parsedConfig)
	if err != nil {
		newRobotCore.handleHardErr(err)
	}
	newRobotCore.clients = newClients

	// Register SocketIO callbacks
	newSignaler, err := signaling.NewRobotSignaler(newRobotCore.clients, newRobotCore.logger, newRobotCore.rtc, parsedConfig)
	if err != nil {
		newRobotCore.handleHardErr(err)
	}
	newRobotCore.signaler = newSignaler

	// Register the core's PeerConnection pointer to be used for WebRTC connection
	newSignaler.RegisterPeerConnection(newRobotCore.rtc)

	return newRobotCore
}

// Start the robot
func (rc *RobotCore) Start() {
	crashed := make(chan bool) // Crash channel to notify main routine when the ROSNode crashes
	go rc.clients.StartROSNode(crashed)

	select {
	case <-crashed:
		rc.logger.Fatalf("ROS Node has stopped running...")
	}
}

// Fatal exceptions helper
func (rc *RobotCore) handleHardErr(err error) {
	rc.logger.Fatalf(err.Error())
}
