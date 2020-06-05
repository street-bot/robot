package signaling

import "encoding/json"

const RobotRegistrationType string = "RReg"

type SignalingMessage interface {
	ToString() string
}

// RobotRegistrationMessage to identify the robot on a new connection
type RobotRegistrationMessage struct {
	Type    string
	Payload RRegPaylod
}

type RRegPaylod struct {
	RobotID string
}

func NewRobotRegistrationMessage(RobotID string) *RobotRegistrationMessage {
	return &RobotRegistrationMessage{
		Type: RobotRegistrationType,
		Payload: RRegPaylod{
			RobotID: RobotID,
		},
	}
}

// ToString serializer
func (r *RobotRegistrationMessage) ToString() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
