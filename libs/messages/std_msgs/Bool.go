// Automatically generated from the message definition "std_msgs/Bool.msg"
// Regenerate with: gengo msg std_msgs/Bool Bool.msg
package std_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

type _MsgBool struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgBool) Text() string {
	return t.text
}

func (t *_MsgBool) Name() string {
	return t.name
}

func (t *_MsgBool) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgBool) NewMessage() ros.Message {
	m := new(Bool)
	m.Data = false
	return m
}

var (
	MsgBool = &_MsgBool{
		`bool data
`,
		"std_msgs/Bool",
		"8b94c1b53db61fb6aed406028ad6332a",
	}
)

type Bool struct {
	Data bool `rosmsg:"data:bool"`
}

func (m *Bool) GetType() ros.MessageType {
	return MsgBool
}

func (m *Bool) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Data)
	return err
}

func (m *Bool) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Data); err != nil {
		return err
	}
	return err
}
