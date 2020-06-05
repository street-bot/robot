FROM ros:melodic-ros-core-bionic

# Install GStreamer dependencies
RUN apt update && \
    apt install -y libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev gstreamer1.0-plugins-good gstreamer1.0-libav \
    ros-melodic-rosserial ros-melodic-rosserial-arduino build-essential

