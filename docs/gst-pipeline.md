# FFMPEG and GStreamer Pipelines!

#### FFMpeg Screen Portion -> /dev/video0
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

To start the pipeline:
```bash
sudo rmmod v4l2loopback
sudo modprobe v4l2loopback
ffmpeg -f x11grab -r 30 -s 640x480 -i :0.0+0,0 -vcodec rawvideo -pix_fmt yuv420p -f v4l2 /dev/video0
```

#### Simple gst pipeline to view the video test source
```bash
gst-launch-1.0 -v videotestsrc ! xvimagesink
```

#### gst pipeline to view screen
```bash
gst-launch-1.0 -v v4l2src device=/dev/video0 ! video/x-raw,format=I420,width=640,height=480 ! queue ! x264enc ! avdec_h264 ! videoconvert ! xvimagesink
```