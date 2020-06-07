package signaling

import "encoding/json"

type SignalingMessage interface {
	ToString() string
}

// Types
const RobotRegistrationType string = "RReg"
const ClientDeregistrationType string = "CDreg"
const OfferType string = "Offer"
const OfferResponseType string = "OfferResponse"

// ------------------------------------------------------------------------------------

// BaseMessage through the WebSocket
type BaseMessage struct {
	Type    string
	Payload interface{}
}

// ToString serializer
func (r *BaseMessage) ToString() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ------------------------------------------------------------------------------------

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

// ------------------------------------------------------------------------------------

// OfferMessage to with SDP string
type OfferMessage struct {
	Type    string
	Payload OfferPayload
}

type OfferPayload struct {
	SDPStr string
}

func NewOfferMessage(SDPStr string) *OfferMessage {
	return &OfferMessage{
		Type: OfferType,
		Payload: OfferPayload{
			SDPStr: SDPStr,
		},
	}
}

// ToString serializer
func (r *OfferMessage) ToString() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ------------------------------------------------------------------------------------

// RobotRegistrationMessage to identify the robot on a new connection
type OfferResponseMessage struct {
	Type    string
	Payload OfferResponsePayload
}

type OfferResponsePayload struct {
	SDPStr string
}

func NewOfferResponseMessage(SDPStr string) *OfferResponseMessage {
	return &OfferResponseMessage{
		Type: OfferResponseType,
		Payload: OfferResponsePayload{
			SDPStr: SDPStr,
		},
	}
}

// ToString serializer
func (r *OfferResponseMessage) ToString() (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// -------------------------------------------------------------------------
// ClientDeregistrationMessage tells the robot that the client is no longer interested in the connection
type ClientDeregistrationMessage struct {
	Type string
}
