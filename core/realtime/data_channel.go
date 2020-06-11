package realtime

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/street-bot/robot/libs/vr2"
	"github.com/street-bot/robot/libs/web"
	"github.com/street-bot/robot/libs/ydlidar_ros_driver"
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

// DataChannelRcvHandler for post-receive actions on DataChannels
func (r *RobotConnection) DataChannelRcvHandler(logger rlog.Logger, config *viper.Viper, dc *webrtc.DataChannel, clients clients.Clients) error {
	// Register DataChannel callbacks to publish to ROS
	controlTopic := "/fromweb"
	lidarTopic := "/laser_fan"

	// Register channel opening handling
	dc.OnOpen(func() {
		logger.Infof("Data channel '%s'-'%d' open", dc.Label(), dc.ID())
		clients.AddROSPub(controlTopic, vr2.MsgVelocity)
		clients.AddROSSub(lidarTopic, ydlidar_ros_driver.MsgLaserFan, lidarMsgCallback(logger, dc))
	})

	// Register channel opening handling
	dc.OnClose(func() {
		logger.Infof("Data channel '%s'-'%d' closed", dc.Label(), dc.ID())
		if err := clients.RemoveROSPub(controlTopic); err != nil {
			logger.Warnf(err.Error())
		}
		if err := clients.RemoveROSSub(lidarTopic); err != nil {
			logger.Warnf(err.Error())
		}
	})

	// Register channel opening handling
	dc.OnError(func(err error) {
		logger.Warnf("Data channel '%s'-'%d' error: %s", dc.Label(), dc.ID(), err.Error())
	})

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		logger.Debugf("Message from DataChannel '%s'-'%d': '%s'", dc.Label(), dc.ID(), string(msg.Data))
		controlMsg, err := web.NewControlMessage(msg.Data)
		if err != nil {
			logger.Errorf("Failed to parse control message: " + err.Error())
			return
		}

		sendMsg := new(vr2.Velocity)
		sendMsg.Forward = int8(controlMsg.Forward)
		sendMsg.Right = int8(controlMsg.Right)
		sendMsg.SpeedLevel = uint8(controlMsg.SpeedLevel)
		clients.ROSPub(controlTopic).Publish(sendMsg)
	})

	return nil
}
