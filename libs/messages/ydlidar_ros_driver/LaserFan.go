// Automatically generated from the message definition "ydlidar_ros_driver/LaserFan.msg"
// Regenerate with: gengo msg ydlidar_ros_driver/LaserFan LaserFan.msg
package ydlidar_ros_driver

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
	"github.com/street-bot/robot/libs/messages/std_msgs"
)

type _MsgLaserFan struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLaserFan) Text() string {
	return t.text
}

func (t *_MsgLaserFan) Name() string {
	return t.name
}

func (t *_MsgLaserFan) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLaserFan) NewMessage() ros.Message {
	m := new(LaserFan)
	m.Header = std_msgs.Header{}
	m.AngleMin = 0.0
	m.AngleMax = 0.0
	m.TimeIncrement = 0.0
	m.ScanTime = 0.0
	m.RangeMin = 0.0
	m.RangeMax = 0.0
	m.Angles = []float32{}
	m.Ranges = []float32{}
	m.Intensities = []float32{}
	return m
}

var (
	MsgLaserFan = &_MsgLaserFan{
		`std_msgs/Header header
float32 angle_min
float32 angle_max
float32 time_increment
float32 scan_time
float32 range_min
float32 range_max
float32[] angles
float32[] ranges
float32[] intensities`,
		"ydlidar_ros_driver/LaserFan",
		"be4554a7f739c3325c744fb261ecf7eb",
	}
)

type LaserFan struct {
	Header        std_msgs.Header `rosmsg:"header:Header"`
	AngleMin      float32         `rosmsg:"angle_min:float32"`
	AngleMax      float32         `rosmsg:"angle_max:float32"`
	TimeIncrement float32         `rosmsg:"time_increment:float32"`
	ScanTime      float32         `rosmsg:"scan_time:float32"`
	RangeMin      float32         `rosmsg:"range_min:float32"`
	RangeMax      float32         `rosmsg:"range_max:float32"`
	Angles        []float32       `rosmsg:"angles:float32[]"`
	Ranges        []float32       `rosmsg:"ranges:float32[]"`
	Intensities   []float32       `rosmsg:"intensities:float32[]"`
}

func (m *LaserFan) GetType() ros.MessageType {
	return MsgLaserFan
}

func (m *LaserFan) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.AngleMin)
	binary.Write(buf, binary.LittleEndian, m.AngleMax)
	binary.Write(buf, binary.LittleEndian, m.TimeIncrement)
	binary.Write(buf, binary.LittleEndian, m.ScanTime)
	binary.Write(buf, binary.LittleEndian, m.RangeMin)
	binary.Write(buf, binary.LittleEndian, m.RangeMax)
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Angles)))
	for _, e := range m.Angles {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Ranges)))
	for _, e := range m.Ranges {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, uint32(len(m.Intensities)))
	for _, e := range m.Intensities {
		binary.Write(buf, binary.LittleEndian, e)
	}
	return err
}

func (m *LaserFan) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleMin); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleMax); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.TimeIncrement); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.ScanTime); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RangeMin); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.RangeMax); err != nil {
		return err
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Angles = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Angles[i]); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Ranges = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Ranges[i]); err != nil {
				return err
			}
		}
	}
	{
		var size uint32
		if err = binary.Read(buf, binary.LittleEndian, &size); err != nil {
			return err
		}
		m.Intensities = make([]float32, int(size))
		for i := 0; i < int(size); i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.Intensities[i]); err != nil {
				return err
			}
		}
	}
	return err
}
