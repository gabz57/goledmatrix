package controller

type KeyboardEventType string

const (
	KeyEventTypeChar   KeyboardEventType = "char"
	KeyEventTypeSymbol KeyboardEventType = "symbol"
)

type KeyboardEventAction int

const (
	PressKey KeyboardEventAction = iota
	HoldKey
	ReleaseKey
)

type KeyboardEvent struct {
	Name   KeyboardEventType   // char, symbol
	Action KeyboardEventAction // press, hold, release
	Data   interface{}
}

func NewKeyboardEvent(name KeyboardEventType, action KeyboardEventAction, data interface{}) *KeyboardEvent {
	return &KeyboardEvent{Name: name, Action: action, Data: data}
}

type KeyboardEventChannel chan *KeyboardEvent

type Keyboard interface {
	Projection() *KeyboardProjection
	Start()
	Stop()
	EventChannel() *KeyboardEventChannel
}

const KeyboardEventChannelSize = 1000

type KeyboardProjection struct {
}

func NewKeyboardProjection() *KeyboardProjection {
	return &KeyboardProjection{}
}
