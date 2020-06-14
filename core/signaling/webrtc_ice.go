package signaling

import (
	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/realtime"
)

const (
	GPSChannelName         = "gps"
	ControlChannelName     = "control"
	LidarChannelName       = "lidar"
	SensorChannelName      = "sensor"
	MiscControlChannelName = "misc-control"
)

// MakeICEStateChangeHandler creates the function to handle ICE connection state changes
func (rs *RobotSignaler) MakeICEStateChangeHandler(rtc realtime.Connection, config *viper.Viper) func(connectionState webrtc.ICEConnectionState) {
	return func(connectionState webrtc.ICEConnectionState) {
		rs.logger.Infof("Connection State has changed %s", connectionState.String())

		if connectionState == webrtc.ICEConnectionStateConnected {
			rs.logger.Debugf("Running PeerConnection OnConnect handler(s)")
			if err := rtc.ICEConnectedPCHandler(rs.logger, config); err != nil {
				rs.logger.Errorf("PeerConnection ICE OnConnect Handler: %s", err.Error())
				return
			}
		} else if connectionState == webrtc.ICEConnectionStateDisconnected {
			rs.logger.Debugf("Running PeerConnection OnDisconnect handler(s)")
			if err := rtc.ICEDisconnectedPCHandler(rs.logger, config); err != nil {
				rs.logger.Errorf("PeerConnection ICE OnDisconnect Handler: %s", err.Error())
				return
			}
		}

		// Other state transition handlers should be added here
	}
}

// MakeDataChannelRcvHandler creates the function to handle DataChannel received events
func (rs *RobotSignaler) MakeDataChannelRcvHandler(rtc realtime.Connection, config *viper.Viper) func(*webrtc.DataChannel) {
	return func(dc *webrtc.DataChannel) {
		rs.logger.Infof("New DataChannel %s %d", dc.Label(), dc.ID())

		var err error
		switch dc.Label() {
		case GPSChannelName:
			err = rtc.GPSChannelRcvHandler(rs.logger, config, dc, rs.clients)
		case ControlChannelName:
			err = rtc.ControlChannelRcvHandler(rs.logger, config, dc, rs.clients)
		case LidarChannelName:
			err = rtc.LidarChannelRcvHandler(rs.logger, config, dc, rs.clients)
		case SensorChannelName:
			err = rtc.SensorChannelRcvHandler(rs.logger, config, dc, rs.clients)
		case MiscControlChannelName:
			err = rtc.MiscControlChannelRcvHandler(rs.logger, config, dc, rs.clients)

		// Other OnDataChannel handlers should be added here

		default:
			rs.logger.Warnf("Opened datachannel %s for unknown purpose!", dc.Label())
		}

		if err != nil {
			rs.logger.Errorf("DataChannel Received Handler: %s", err.Error())
			return
		}

	}
}
