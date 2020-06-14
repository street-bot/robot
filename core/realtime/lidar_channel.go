package realtime

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/street-bot/robot/libs/messages/ydlidar_ros_driver"
)

func lidarMsgCallback(logger rlog.Logger, dc *webrtc.DataChannel) func(message *ydlidar_ros_driver.LaserFan) {
	return func(message *ydlidar_ros_driver.LaserFan) {
		msg, err := json.Marshal(message)
		if err != nil {
			logger.Errorf("Unmarshal LiDAR message: %s", err.Error())
		}
		dc.SendText(string(msg))
	}
}

// LidarChannelRcvHandler for post-receive actions on DataChannels
func (r *RobotConnection) LidarChannelRcvHandler(logger rlog.Logger, config *viper.Viper, dc *webrtc.DataChannel, clients clients.Clients) error {
	// Register DataChannel callbacks to publish to ROS
	topic := "/laser_fan"

	// Register channel opening handling
	dc.OnOpen(func() {
		logger.Infof("Data channel '%s'-'%d' open", dc.Label(), dc.ID())
		clients.AddROSSub(topic, ydlidar_ros_driver.MsgLaserFan, lidarMsgCallback(logger, dc))
	})

	// Register channel opening handling
	dc.OnClose(func() {
		logger.Infof("Data channel '%s'-'%d' closed", dc.Label(), dc.ID())
		if err := clients.RemoveROSSub(topic); err != nil {
			logger.Warnf(err.Error())
		}
	})

	// Register channel opening handling
	dc.OnError(func(err error) {
		logger.Errorf("Data channel '%s'-'%d' error: %s", dc.Label(), dc.ID(), err.Error())
	})

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		logger.Warnf("Unexpected message from DataChannel '%s'-'%d': '%s'", dc.Label(), dc.ID(), string(msg.Data))
	})

	return nil
}
