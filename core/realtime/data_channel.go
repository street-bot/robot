package realtime

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/street-bot/robot/libs/messages/sensor_msgs"
	"github.com/street-bot/robot/libs/messages/vr2"
	"github.com/street-bot/robot/libs/messages/ydlidar_ros_driver"
	"github.com/street-bot/robot/libs/web"
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

func gpsMsgCallback(logger rlog.Logger, dc *webrtc.DataChannel) func(message *sensor_msgs.NavSatFix) {
	return func(message *sensor_msgs.NavSatFix) {
		msg, err := json.Marshal(message)
		if err != nil {
			logger.Errorf("Unmarshal GPS message: %s", err.Error())
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

// ControlChannelRcvHandler for post-receive actions on datachannel for control signals
func (r *RobotConnection) ControlChannelRcvHandler(logger rlog.Logger, config *viper.Viper, dc *webrtc.DataChannel, clients clients.Clients) error {
	// Register DataChannel callbacks to publish to ROS
	topic := "/fromweb"

	// Register channel opening handling
	dc.OnOpen(func() {
		logger.Infof("Data channel '%s'-'%d' open", dc.Label(), dc.ID())
		clients.AddROSPub(topic, vr2.MsgVelocity)
	})

	// Register channel opening handling
	dc.OnClose(func() {
		logger.Infof("Data channel '%s'-'%d' closed", dc.Label(), dc.ID())
		if err := clients.RemoveROSPub(topic); err != nil {
			logger.Warnf(err.Error())
		}
	})

	// Register channel opening handling
	dc.OnError(func(err error) {
		logger.Errorf("Data channel '%s'-'%d' error: %s", dc.Label(), dc.ID(), err.Error())
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
		clients.ROSPub(topic).Publish(sendMsg)
	})

	return nil
}

// GPSChannelRcvHandler for post-receive actions on GPS data channel
func (r *RobotConnection) GPSChannelRcvHandler(logger rlog.Logger, config *viper.Viper, dc *webrtc.DataChannel, clients clients.Clients) error {
	// Register DataChannel callbacks to publish to ROS
	topic := "/gps"

	// Register channel opening handling
	dc.OnOpen(func() {
		logger.Infof("Data channel '%s'-'%d' open", dc.Label(), dc.ID())
		clients.AddROSSub(topic, sensor_msgs.MsgNavSatFix, gpsMsgCallback(logger, dc))
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
