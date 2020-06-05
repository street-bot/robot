package realtime

import (
	"fmt"

	"github.com/pion/webrtc/v2"
	"github.com/spf13/viper"
	gst "github.com/street-bot/robot/libs/gstreamer-src"
	rlog "github.com/street-bot/robot/libs/log"
)

// ICEConnectedPCHandler for post-connection actions on PeerConnections
func (r *RobotConnection) ICEConnectedPCHandler(logger rlog.Logger, config *viper.Viper) error {
	// Grab the video device
	vidDevice := config.GetString("multimedia.camera.device")
	if vidDevice == "" {
		return fmt.Errorf("configuration multimedia.camera.device not specified")
	}

	// Override format string if specified in config
	formatString := "video/x-h264,width=640,height=480,framerate=24/1"
	overrideFormatString := config.GetString("multimedia.camera.format")
	if overrideFormatString != "" {
		formatString = overrideFormatString
	}
	// Video source string
	videoSrc := "v4l2src device=" + vidDevice + " ! " + formatString + " ! avdec_h264 ! videoconvert" // Logitech C920

	// Get track configuration
	trackConfigs := config.GetStringMapString("multimedia.tracks")
	cameraTrackName := config.GetString("multimedia.camera.track")
	if cameraTrackName == "" {
		return fmt.Errorf("configuration multimedia.camera.track not specified")
	}
	payloadType, ok := trackConfigs[cameraTrackName]
	if !ok {
		return fmt.Errorf("track %s not found in multimediea.tracks", cameraTrackName)
	}

	// Start pushing buffers on the track(s)
	newPipeline := gst.CreatePipeline(payloadType, []*webrtc.Track{r.Track("video")}, videoSrc)
	if err := newPipeline.Start(); err != nil {
		logger.Warnf("gst create pipeline: ", err.Error())
		return nil
	}
	r.pipelines[cameraTrackName] = newPipeline

	return nil
}

// ICEDisconnectedPCHandler for post-connection actions on PeerConnections
func (r *RobotConnection) ICEDisconnectedPCHandler(logger rlog.Logger, config *viper.Viper) error {
	// Get the target video pipeline
	cameraTrackName := config.GetString("multimedia.camera.track")
	if cameraTrackName == "" {
		return fmt.Errorf("configuration multimedia.camera.track not specified")
	}
	logger.Infof("Attempting to stop pipeline %s", cameraTrackName)

	pipeline, ok := r.pipelines[cameraTrackName]
	if !ok {
		return fmt.Errorf("pipeline %s not found", cameraTrackName)
	}

	if err := pipeline.Stop(); err != nil {
		logger.Warnf("gst create pipeline: ", err.Error())
		return nil
	}
	logger.Infof("Pipeline %s has stopped", cameraTrackName)

	return nil
}
