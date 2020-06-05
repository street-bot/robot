// Automatically generated from the message definition "vr2/Velocity.msg"
// To regenerate: gengo msg vr2/Velocity Velocity.msg
package vr2

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgVelocity struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgVelocity) Text() string {
	return t.text
}

func (t *_MsgVelocity) Name() string {
	return t.name
}

func (t *_MsgVelocity) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgVelocity) NewMessage() ros.Message {
	m := new(Velocity)
	m.Forward = 0
	m.Right = 0
	m.SpeedLevel = 0
	return m
}

var (
	MsgVelocity = &_MsgVelocity{
		`int8 forward
int8 right
uint8 speedLevel`,
		"vr2/Velocity",
		"17c0b027db97ab2c68fd79cc82e4a77d",
	}
)

type Velocity struct {
	Forward    int8  `rosmsg:"forward:int8"`
	Right      int8  `rosmsg:"right:int8"`
	SpeedLevel uint8 `rosmsg:"speedLevel:uint8"`
}

func (m *Velocity) GetType() ros.MessageType {
	return MsgVelocity
}

func (m *Velocity) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Forward)
	binary.Write(buf, binary.LittleEndian, m.Right)
	binary.Write(buf, binary.LittleEndian, m.SpeedLevel)
	return err
}

func (m *Velocity) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Forward); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Right); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.SpeedLevel); err != nil {
		return err
	}
	return err
}
