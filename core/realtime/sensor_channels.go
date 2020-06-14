package realtime

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/street-bot/robot/libs/messages/sensor_msgs"
	"github.com/street-bot/robot/libs/messages/std_msgs"
)

// Anything that pushes an std_msgs/bool type
func boolMsgCallback(logger rlog.Logger, dc *webrtc.DataChannel) func(message *std_msgs.Bool) {
	return func(message *std_msgs.Bool) {
		wrappedMsg := NewDataChannelMessage(FoodBoxLatchState, message)
		msg, err := json.Marshal(wrappedMsg)
		if err != nil {
			logger.Errorf("Unmarshal BatteryState message: %s", err.Error())
		}
		dc.SendText(string(msg))
	}
}

func batteryStateMsgCallback(logger rlog.Logger, dc *webrtc.DataChannel) func(message *sensor_msgs.BatteryState) {
	return func(message *sensor_msgs.BatteryState) {
		wrappedMsg := NewDataChannelMessage(BatteryState, message)
		msg, err := json.Marshal(wrappedMsg)
		if err != nil {
			logger.Errorf("Unmarshal BatteryState message: %s", err.Error())
		}
		dc.SendText(string(msg))
	}
}

func temperatureMsgCallback(logger rlog.Logger, dc *webrtc.DataChannel, Type string) func(message *sensor_msgs.Temperature) {
	return func(message *sensor_msgs.Temperature) {
		wrappedMsg := NewDataChannelMessage(Type, message)
		msg, err := json.Marshal(wrappedMsg)
		if err != nil {
			logger.Errorf("Unmarshal Temperature message: %s", err.Error())
		}
		dc.SendText(string(msg))
	}
}

// SensorChannelRcvHandler for post-receive actions on sensor data channel
func (r *RobotConnection) SensorChannelRcvHandler(logger rlog.Logger, config *viper.Viper, dc *webrtc.DataChannel, clients clients.Clients) error {
	// Register DataChannel callbacks to publish to ROS
	batteryTopic := "/sensor_msgs/BatteryState"
	controlBoxTempTopic := "/sensor_msgs/ControlBoxTemerature"
	foodBoxTempTopic := "/sensor_msgs/FoodBoxTemerature"
	foodBoxLatchStateTopic := "/std_msgs/BoxClosed"

	// Register channel opening handling
	dc.OnOpen(func() {
		logger.Infof("Data channel '%s'-'%d' open", dc.Label(), dc.ID())
		clients.AddROSSub(batteryTopic, sensor_msgs.MsgBatteryState, batteryStateMsgCallback(logger, dc))
		clients.AddROSSub(controlBoxTempTopic, sensor_msgs.MsgTemperature, temperatureMsgCallback(logger, dc, ControlBoxState))
		clients.AddROSSub(foodBoxTempTopic, sensor_msgs.MsgTemperature, temperatureMsgCallback(logger, dc, FoodBoxState))
		clients.AddROSSub(foodBoxLatchStateTopic, std_msgs.MsgBool, boolMsgCallback(logger, dc))
	})

	// Register channel opening handling
	dc.OnClose(func() {
		logger.Infof("Data channel '%s'-'%d' closed", dc.Label(), dc.ID())
		if err := clients.RemoveROSSub(batteryTopic); err != nil {
			logger.Warnf(err.Error())
		}
		if err := clients.RemoveROSSub(controlBoxTempTopic); err != nil {
			logger.Warnf(err.Error())
		}
		if err := clients.RemoveROSSub(foodBoxTempTopic); err != nil {
			logger.Warnf(err.Error())
		}
		if err := clients.RemoveROSSub(foodBoxLatchStateTopic); err != nil {
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
