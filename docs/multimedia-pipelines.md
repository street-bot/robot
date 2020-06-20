# Multimedia Pipelines!

## Overview
Proper pipeline setup enables stream splitting of a single camera source to multiple consumers. This is very useful for applications where the same camera stream should be "Teed" to feed the WebRTC sink and ROS video node sink. Stream duplication helps isolate pipeline failures and needs to be done at the host OS level since it requires the `v4l2loopback` kernel module.

## Setup
#### Properly setting up `v4l2loopback`
The `v4l2loopback` kernel has a known bug. To Fix:
```bash
sudo apt-get remove -y v4l2loopback-dkms
sudo apt-get install -y build-essential libelf-dev linux-headers-$(uname -r) unzip
wget https://github.com/umlaeute/v4l2loopback/archive/master.zip
unzip master.zip
cd v4l2loopback-master
make
sudo make install
sudo depmod -a
sudo modprobe v4l2loopback
```

#### FFMPEG duplicate one video device into multiple copies
In order to initiate the pipeline successfully, the origin input `video0` must be present and streaming data.
It's possible to "open" the video stream on the same device with multiple sinks. However, if one of the sinks on a given device closes, it will cause the entire pipeline to crash. Duplicating the devices and opening the duplicated device files isolates the pipelines thereby reducing blast radius.
```bash
sudo modprobe v4l2loopback video_nr=0,1,2
ffmpeg -i /dev/video0 -codec copy -f v4l2 /dev/video1 -codec copy -f v4l2 /dev/video2
```

#### FFMpeg Screen Portion -> /dev/video0
```bash
sudo rmmod v4l2loopback
sudo modprobe v4l2loopback # video_nr=0,1 # <- Enumeration of the virual v4l2 devices to create
ffmpeg -f x11grab -r 30 -s 640x480 -i :0.0+0,0 -vcodec rawvideo -pix_fmt yuv420p -f v4l2 /dev/video0
```

## GStreamer Pipelines
#### Simple gst pipeline to view the video test source
```bash
gst-launch-1.0 -v videotestsrc ! xvimagesink
```

#### gst pipeline to view screen through x264 encoding and decoding
```bash
gst-launch-1.0 -v v4l2src device=/dev/video1 ! video/x-raw,format=I420,width=640,height=480 ! queue ! x264enc ! avdec_h264 ! videoconvert ! xvimagesink
```

#### Split gst pipeline with tee
```bash
gst-launch-1.0 -v v4l2src device=/dev/video0 ! tee name=t \
  t. ! queue ! video/x-raw,format=I420,width=640,height=480 ! videoconvert ! xvimagesink \
  t. ! queue ! v4l2sink device=/dev/video1

```