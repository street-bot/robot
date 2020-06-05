package web

import "encoding/json"

// ControlMessage from web-based clients
type ControlMessage struct {
	Forward    int
	Right      int
	SpeedLevel uint
}

// NewControlMessage constructor from a given JSON string
func NewControlMessage(input []byte) (*ControlMessage, error) {
	var newControlMessage ControlMessage
	if err := json.Unmarshal(input, &newControlMessage); err != nil {
		return nil, err
	}
	return &newControlMessage, nil
}
