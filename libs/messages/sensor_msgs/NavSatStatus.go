// Automatically generated from the message definition "sensor_msgs/NavSatStatus.msg"
// Regenerate with: gengo msg sensor_msgs/NavSatStatus NavSatStatus.msg
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	NavSatStatus_STATUS_NO_FIX   int8   = -1
	NavSatStatus_STATUS_FIX      int8   = 0
	NavSatStatus_STATUS_SBAS_FIX int8   = 1
	NavSatStatus_STATUS_GBAS_FIX int8   = 2
	NavSatStatus_SERVICE_GPS     uint16 = 1
	NavSatStatus_SERVICE_GLONASS uint16 = 2
	NavSatStatus_SERVICE_COMPASS uint16 = 4
	NavSatStatus_SERVICE_GALILEO uint16 = 8
)

type _MsgNavSatStatus struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgNavSatStatus) Text() string {
	return t.text
}

func (t *_MsgNavSatStatus) Name() string {
	return t.name
}

func (t *_MsgNavSatStatus) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgNavSatStatus) NewMessage() ros.Message {
	m := new(NavSatStatus)
	m.Status = 0
	m.Service = 0
	return m
}

var (
	MsgNavSatStatus = &_MsgNavSatStatus{
		`# Navigation Satellite fix status for any Global Navigation Satellite System

# Whether to output an augmented fix is determined by both the fix
# type and the last time differential corrections were received.  A
# fix is valid when status >= STATUS_FIX.

int8 STATUS_NO_FIX =  -1        # unable to fix position
int8 STATUS_FIX =      0        # unaugmented fix
int8 STATUS_SBAS_FIX = 1        # with satellite-based augmentation
int8 STATUS_GBAS_FIX = 2        # with ground-based augmentation

int8 status

# Bits defining which Global Navigation Satellite System signals were
# used by the receiver.

uint16 SERVICE_GPS =     1
uint16 SERVICE_GLONASS = 2
uint16 SERVICE_COMPASS = 4      # includes BeiDou.
uint16 SERVICE_GALILEO = 8

uint16 service`,
		"sensor_msgs/NavSatStatus",
		"331cdbddfa4bc96ffc3b9ad98900a54c",
	}
)

type NavSatStatus struct {
	Status  int8   `rosmsg:"status:int8"`
	Service uint16 `rosmsg:"service:uint16"`
}

func (m *NavSatStatus) GetType() ros.MessageType {
	return MsgNavSatStatus
}

func (m *NavSatStatus) Serialize(buf *bytes.Buffer) error {
	var err error
	binary.Write(buf, binary.LittleEndian, m.Status)
	binary.Write(buf, binary.LittleEndian, m.Service)
	return err
}

func (m *NavSatStatus) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = binary.Read(buf, binary.LittleEndian, &m.Status); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Service); err != nil {
		return err
	}
	return err
}
