package signaling

import (
	"github.com/pion/webrtc/v2"
	"github.com/street-bot/robot/core/realtime"
	"github.com/street-bot/robot/libs/signal"
)

// RespondOffer for WebRTC
func (rs *RobotSignaler) RespondOffer(pc *webrtc.PeerConnection, offerStr string) (string, error) {
	// Wait for the offer to be pasted
	offer := webrtc.SessionDescription{}

	var err error
	err = signal.Decode(offerStr, &offer)
	if err != nil {
		return "", err
	}

	// Set the remote SessionDescription
	err = pc.SetRemoteDescription(offer)
	if err != nil {
		return "", err
	}

	// Create an answer
	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		return "", err
	}

	// Sets the LocalDescription, and starts our UDP listeners
	err = pc.SetLocalDescription(answer)
	if err != nil {
		return "", err
	}

	answerStr, err := signal.Encode(answer)
	if err != nil {
		return "", err
	}
	return answerStr, nil
}

// OnOffer event handler constructor
func (rs *RobotSignaler) onOffer(rtc realtime.Connection) func(string) {
	return func(msg string) {
		offerStr := msg // In case the msg is not directly the offer
		rs.logger.Infof("Received offer")
		rs.logger.Debugf(offerStr)

		// Create realtime connection object
		// Edit below stanza to add more data channels and tracks
		tracks := rs.config.GetStringMapString("multimedia.tracks")
		newRtc, err := realtime.NewRobotConnection(tracks)
		if err != nil {
			rs.logger.Errorf("WebRTC creation error: %s", err.Error())
			return
		}
		rtc = newRtc

		// Set the handler for ICE connection state change
		rtc.PeerConnection().OnICEConnectionStateChange(rs.MakeICEStateChangeHandler(rtc, rs.config))

		// Set the handler for DataChannel received
		rtc.PeerConnection().OnDataChannel(rs.MakeDataChannelRcvHandler(rtc, rs.config))

		responseStr, err := rs.RespondOffer(rtc.PeerConnection(), offerStr)
		if err != nil {
			rs.logger.Errorf("WebRTC offer response error: %s", err.Error())
			return
		}

		// Send response
		// rs.clients.SocketIO().Emit("/offerResponse", responseStr)	// TODO
		rs.logger.Infof("Sent response")
		rs.logger.Debugf(responseStr)
	}
}
