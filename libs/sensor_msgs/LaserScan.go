// Automatically generated from the message definition "sensor_msgs/LaserScan.msg"
// To regenerate: gengo msg sensor_msgs/LaserScan LaserScan.msg
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/street-bot/robot/libs/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgLaserScan struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLaserScan) Text() string {
	return t.text
}

func (t *_MsgLaserScan) Name() string {
	return t.name
}

func (t *_MsgLaserScan) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLaserScan) NewMessage() ros.Message {
	m := new(LaserScan)
	m.Header = std_msgs.Header{}
	m.AngleMin = 0.0
	m.AngleMax = 0.0
	m.AngleIncrement = 0.0
	m.TimeIncrement = 0.0
	m.ScanTime = 0.0
	m.RangeMin = 0.0
	m.RangeMax = 0.0
	m.Ranges = []float32{}
	m.Intensities = []float32{}
	return m
}

var (
	MsgLaserScan = &_MsgLaserScan{
		`std_msgs/Header header
float32 angle_min
float32 angle_max
float32 angle_increment
float32 time_increment
float32 scan_time
float32 range_min
float32 range_max
float32[] ranges
float32[] intensities`,
		"sensor_msgs/LaserScan",
		"90c7ef2dc6895d81024acba2ac42f369",
	}
)

type LaserScan struct {
	Header         std_msgs.Header `rosmsg:"header:Header"`
	AngleMin       float32         `rosmsg:"angle_min:float32"`
	AngleMax       float32         `rosmsg:"angle_max:float32"`
	AngleIncrement float32         `rosmsg:"angle_increment:float32"`
	TimeIncrement  float32         `rosmsg:"time_increment:float32"`
	ScanTime       float32         `rosmsg:"scan_time:float32"`
	RangeMin       float32         `rosmsg:"range_min:float32"`
	RangeMax       float32         `rosmsg:"range_max:float32"`
	Ranges         []float32       `rosmsg:"ranges:float32[]"`
	Intensities    []float32       `rosmsg:"intensities:float32[]"`
}

func (m *LaserScan) GetType() ros.MessageType {
	return MsgLaserScan
}

func (m *LaserScan) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.AngleMin)
	binary.Write(buf, binary.LittleEndian, m.AngleMax)
	binary.Write(buf, binary.LittleEndian, m.AngleIncrement)
	binary.Write(buf, binary.LittleEndian, m.TimeIncrement)
	binary.Write(buf, binary.LittleEndian, m.ScanTime)
	binary.Write(buf, binary.LittleEndian, m.RangeMin)
	binary.Write(buf, binary.LittleEndian, m.RangeMax)
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

func (m *LaserScan) Deserialize(buf *bytes.Reader) error {
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
	if err = binary.Read(buf, binary.LittleEndian, &m.AngleIncrement); err != nil {
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
