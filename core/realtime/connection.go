package realtime

import (
	"fmt"
	"math/rand"

	"github.com/street-bot/robot/core/clients"
	gst "github.com/street-bot/robot/libs/gstreamer-src"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	rlog "github.com/street-bot/robot/libs/log"
)

// Connection interface abstracts WebRTC realtime constructs
type Connection interface {
	Track(string) *webrtc.Track
	DataChannel(string) *webrtc.DataChannel
	PeerConnection() *webrtc.PeerConnection

	RegisterDataChannel(string, *webrtc.DataChannel)

	// PeerConnection ICE state change hooks
	ICEConnectedPCHandler(rlog.Logger, *viper.Viper) error
	ICEDisconnectedPCHandler(rlog.Logger, *viper.Viper) error

	// Datachannel received handlers
	ControlChannelRcvHandler(rlog.Logger, *viper.Viper, *webrtc.DataChannel, clients.Clients) error
	GPSChannelRcvHandler(rlog.Logger, *viper.Viper, *webrtc.DataChannel, clients.Clients) error
	LidarChannelRcvHandler(rlog.Logger, *viper.Viper, *webrtc.DataChannel, clients.Clients) error
	SensorChannelRcvHandler(rlog.Logger, *viper.Viper, *webrtc.DataChannel, clients.Clients) error
	MiscControlChannelRcvHandler(rlog.Logger, *viper.Viper, *webrtc.DataChannel, clients.Clients) error
}

// RobotConnection holds the robot's realtime connection objects
type RobotConnection struct {
	pc           *webrtc.PeerConnection
	tracks       map[string]*webrtc.Track
	pipelines    map[string]*gst.Pipeline // Pipeline keys should match up with track keys
	dataChannels map[string]*webrtc.DataChannel
}

// NewRobotConnection with the given Tracks and DataChannels
//
// User can then call the Track() and DataChannel() accessors to get a specific track/channel
func NewRobotConnection(tracks map[string]string, iceServers []webrtc.ICEServer) (Connection, error) {
	newRTC := new(RobotConnection)
	newRTC.tracks = make(map[string]*webrtc.Track)
	newRTC.pipelines = make(map[string]*gst.Pipeline)
	newRTC.dataChannels = make(map[string]*webrtc.DataChannel)

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: iceServers,
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}
	newRTC.pc = peerConnection

	// Create a multi-media tracks
	for k, v := range tracks {
		// Key is the track id
		// Value is the track payload type
		payloadType, err := mapPayloadStringCode(v)
		if err != nil {
			return nil, err
		}
		newTrack, err := newRTC.pc.NewTrack(payloadType, rand.Uint32(), k, k) // id and label same name
		if err != nil {
			return nil, err
		}
		_, err = newRTC.pc.AddTrack(newTrack)
		if err != nil {
			return nil, err
		}
		newRTC.tracks[k] = newTrack
	}

	return newRTC, nil
}

// PeerConnection accessor
func (r *RobotConnection) PeerConnection() *webrtc.PeerConnection {
	return r.pc
}

// DataChannel accessor
func (r *RobotConnection) DataChannel(name string) *webrtc.DataChannel {
	return r.dataChannels[name]
}

// Track accessor
func (r *RobotConnection) Track(name string) *webrtc.Track {
	return r.tracks[name]
}

// RegisterDataChannel setter
func (r *RobotConnection) RegisterDataChannel(name string, dc *webrtc.DataChannel) {
	r.dataChannels[name] = dc
}

// Map the string payload type into the payload code
func mapPayloadStringCode(input string) (uint8, error) {
	switch input {
	case webrtc.PCMU:
		return webrtc.DefaultPayloadTypePCMU, nil
	case webrtc.PCMA:
		return webrtc.DefaultPayloadTypePCMA, nil
	case webrtc.G722:
		return webrtc.DefaultPayloadTypeG722, nil
	case webrtc.Opus:
		return webrtc.DefaultPayloadTypeOpus, nil
	case webrtc.VP8:
		return webrtc.DefaultPayloadTypeVP8, nil
	case webrtc.VP9:
		return webrtc.DefaultPayloadTypeVP9, nil
	case webrtc.H264:
		return webrtc.DefaultPayloadTypeH264, nil
	default:
		return 255, fmt.Errorf("Unsupported payload type: %s", input)
	}
}
