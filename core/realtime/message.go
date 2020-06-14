package realtime

// DataChannelMessage wrapper
type DataChannelMessage struct {
	Type string
	Msg  interface{}
}

// NewDataChannelMessage constructor
func NewDataChannelMessage(Type string, Msg interface{}) DataChannelMessage {
	return DataChannelMessage{
		Type: Type,
		Msg:  Msg,
	}
}

const (
	BatteryStateMsg = "BatteryStateMsg"
	FoodBoxTemp     = "FoodBoxTemp"
	ControlBoxTemp  = "ControlBoxTemp"
)
