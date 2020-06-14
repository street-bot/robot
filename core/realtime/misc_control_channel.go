package realtime

import (
	"encoding/json"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	"github.com/street-bot/robot/core/clients"
	rlog "github.com/street-bot/robot/libs/log"
	"github.com/street-bot/robot/libs/messages/std_msgs"
)

// MiscControlChannelRcvHandler for post-receive actions on datachannel for misc control signals
func (r *RobotConnection) MiscControlChannelRcvHandler(logger rlog.Logger, config *viper.Viper, dc *webrtc.DataChannel, clients clients.Clients) error {
	// Register DataChannel callbacks to publish to ROS
	boxLatchTopic := "/latch"

	// Register channel opening handling
	dc.OnOpen(func() {
		logger.Infof("Data channel '%s'-'%d' open", dc.Label(), dc.ID())
		clients.AddROSPub(boxLatchTopic, std_msgs.MsgBool)
	})

	// Register channel opening handling
	dc.OnClose(func() {
		logger.Infof("Data channel '%s'-'%d' closed", dc.Label(), dc.ID())
		if err := clients.RemoveROSPub(boxLatchTopic); err != nil {
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
		var parsedMsg DataChannelMessage

		if err := json.Unmarshal(msg.Data, &parsedMsg); err != nil {
			logger.Errorf("Failed to parse misc control message: " + err.Error())
			return
		}

		switch parsedMsg.Type {
		case BoxLatchControl:
			sendMsg := new(std_msgs.Bool)
			sendMsg.Data = parsedMsg.Msg.(bool)
			clients.ROSPub(boxLatchTopic).Publish(sendMsg)
		default:
			logger.Warnf("Unsupported misc control message type %s", parsedMsg.Type)
		}
	})

	return nil
}
