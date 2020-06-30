FROM ros:melodic-ros-core-bionic

# Install GStreamer dependencies
RUN apt update && \
    apt install -y build-essential \
        libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev gstreamer1.0-plugins-good gstreamer1.0-libav gstreamer1.0-plugins-ugly\
        ros-melodic-rosserial ros-melodic-rosserial-arduino ros-melodic-video-stream-opencv

