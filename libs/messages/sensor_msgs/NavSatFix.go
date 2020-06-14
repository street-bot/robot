// Automatically generated from the message definition "sensor_msgs/NavSatFix.msg"
// Regenerate with: gengo msg sensor_msgs/NavSatFix NavSatFix.msg
package sensor_msgs

import (
	"bytes"
	"encoding/binary"

	"github.com/street-bot/robot/libs/messages/std_msgs"

	"github.com/fetchrobotics/rosgo/ros"
)

const (
	NavSatFix_COVARIANCE_TYPE_UNKNOWN        uint8 = 0
	NavSatFix_COVARIANCE_TYPE_APPROXIMATED   uint8 = 1
	NavSatFix_COVARIANCE_TYPE_DIAGONAL_KNOWN uint8 = 2
	NavSatFix_COVARIANCE_TYPE_KNOWN          uint8 = 3
)

type _MsgNavSatFix struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgNavSatFix) Text() string {
	return t.text
}

func (t *_MsgNavSatFix) Name() string {
	return t.name
}

func (t *_MsgNavSatFix) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgNavSatFix) NewMessage() ros.Message {
	m := new(NavSatFix)
	m.Header = std_msgs.Header{}
	m.Status = NavSatStatus{}
	m.Latitude = 0.0
	m.Longitude = 0.0
	m.Altitude = 0.0
	for i := 0; i < 9; i++ {
		m.PositionCovariance[i] = 0.0
	}
	m.PositionCovarianceType = 0
	return m
}

var (
	MsgNavSatFix = &_MsgNavSatFix{
		`# Navigation Satellite fix for any Global Navigation Satellite System
#
# Specified using the WGS 84 reference ellipsoid

# header.stamp specifies the ROS time for this measurement (the
#        corresponding satellite time may be reported using the
#        sensor_msgs/TimeReference message).
#
# header.frame_id is the frame of reference reported by the satellite
#        receiver, usually the location of the antenna.  This is a
#        Euclidean frame relative to the vehicle, not a reference
#        ellipsoid.
std_msgs/Header header

# satellite fix status information
sensor_msgs/NavSatStatus status

# Latitude [degrees]. Positive is north of equator; negative is south.
float64 latitude

# Longitude [degrees]. Positive is east of prime meridian; negative is west.
float64 longitude

# Altitude [m]. Positive is above the WGS 84 ellipsoid
# (quiet NaN if no altitude is available).
float64 altitude

# Position covariance [m^2] defined relative to a tangential plane
# through the reported position. The components are East, North, and
# Up (ENU), in row-major order.
#
# Beware: this coordinate system exhibits singularities at the poles.

float64[9] position_covariance

# If the covariance of the fix is known, fill it in completely. If the
# GPS receiver provides the variance of each measurement, put them
# along the diagonal. If only Dilution of Precision is available,
# estimate an approximate covariance from that.

uint8 COVARIANCE_TYPE_UNKNOWN = 0
uint8 COVARIANCE_TYPE_APPROXIMATED = 1
uint8 COVARIANCE_TYPE_DIAGONAL_KNOWN = 2
uint8 COVARIANCE_TYPE_KNOWN = 3

uint8 position_covariance_type
`,
		"sensor_msgs/NavSatFix",
		"2d3a8cd499b9b4a0249fb98fd05cfa48",
	}
)

type NavSatFix struct {
	Header                 std_msgs.Header `rosmsg:"header:Header"`
	Status                 NavSatStatus    `rosmsg:"status:NavSatStatus"`
	Latitude               float64         `rosmsg:"latitude:float64"`
	Longitude              float64         `rosmsg:"longitude:float64"`
	Altitude               float64         `rosmsg:"altitude:float64"`
	PositionCovariance     [9]float64      `rosmsg:"position_covariance:float64[9]"`
	PositionCovarianceType uint8           `rosmsg:"position_covariance_type:uint8"`
}

func (m *NavSatFix) GetType() ros.MessageType {
	return MsgNavSatFix
}

func (m *NavSatFix) Serialize(buf *bytes.Buffer) error {
	var err error
	if err = m.Header.Serialize(buf); err != nil {
		return err
	}
	if err = m.Status.Serialize(buf); err != nil {
		return err
	}
	binary.Write(buf, binary.LittleEndian, m.Latitude)
	binary.Write(buf, binary.LittleEndian, m.Longitude)
	binary.Write(buf, binary.LittleEndian, m.Altitude)
	for _, e := range m.PositionCovariance {
		binary.Write(buf, binary.LittleEndian, e)
	}
	binary.Write(buf, binary.LittleEndian, m.PositionCovarianceType)
	return err
}

func (m *NavSatFix) Deserialize(buf *bytes.Reader) error {
	var err error = nil
	if err = m.Header.Deserialize(buf); err != nil {
		return err
	}
	if err = m.Status.Deserialize(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Latitude); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Longitude); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.Altitude); err != nil {
		return err
	}
	{
		for i := 0; i < 9; i++ {
			if err = binary.Read(buf, binary.LittleEndian, &m.PositionCovariance[i]); err != nil {
				return err
			}
		}
	}
	if err = binary.Read(buf, binary.LittleEndian, &m.PositionCovarianceType); err != nil {
		return err
	}
	return err
}
