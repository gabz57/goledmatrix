package controller

import (
	"github.com/kpeu3i/gods4"
)

type GamepadEventChannel chan *GamepadEvent

type GamepadEvent struct {
	Name   string // cross, gyroscope
	Action string // press, release, swipe, move, update
	Data   interface{}
}

type Stick struct {
	X byte
	Y byte
}

type Touchpad struct {
	Press bool
	Swipe []Touch
}

type Touch struct {
	IsActive bool
	X        byte
	Y        byte
}

type Accelerometer struct {
	X int16
	Y int16
	Z int16
}

type Gyroscope struct {
	Roll  int16
	Yaw   int16
	Pitch int16
}

type Battery struct {
	Capacity         byte
	IsCharging       bool
	IsCableConnected bool
}

var allEvents = []gods4.Event{
	gods4.EventCrossPress,
	gods4.EventCrossRelease,
	gods4.EventCirclePress,
	gods4.EventCircleRelease,
	gods4.EventSquarePress,
	gods4.EventSquareRelease,
	gods4.EventTrianglePress,
	gods4.EventTriangleRelease,
	gods4.EventL1Press,
	gods4.EventL1Release,
	gods4.EventL2Press,
	gods4.EventL2Release,
	gods4.EventL3Press,
	gods4.EventL3Release,
	gods4.EventR1Press,
	gods4.EventR1Release,
	gods4.EventR2Press,
	gods4.EventR2Release,
	gods4.EventR3Press,
	gods4.EventR3Release,
	gods4.EventDPadUpPress,
	gods4.EventDPadUpRelease,
	gods4.EventDPadDownPress,
	gods4.EventDPadDownRelease,
	gods4.EventDPadLeftPress,
	gods4.EventDPadLeftRelease,
	gods4.EventDPadRightPress,
	gods4.EventDPadRightRelease,
	gods4.EventSharePress,
	gods4.EventShareRelease,
	gods4.EventOptionsPress,
	gods4.EventOptionsRelease,
	gods4.EventTouchpadSwipe,
	gods4.EventTouchpadPress,
	gods4.EventTouchpadRelease,
	gods4.EventPSPress,
	gods4.EventPSRelease,
	gods4.EventLeftStickMove,
	gods4.EventRightStickMove,
	gods4.EventAccelerometerUpdate,
	gods4.EventGyroscopeUpdate,
	gods4.EventBatteryUpdate,
}

//var byteEvents = []gods4.Event{
//	gods4.EventL2Press,
//	gods4.EventL2Release,
//	gods4.EventR2Press,
//	gods4.EventR2Release,
//}
//var touchpadEvents = []gods4.Event{
//	gods4.EventTouchpadSwipe,
//	gods4.EventTouchpadPress,
//	gods4.EventTouchpadRelease,
//}
//var stickEvents = []gods4.Event{
//	gods4.EventLeftStickMove,
//	gods4.EventRightStickMove,
//}
//var accelerometerEvents = []gods4.Event{
//	gods4.EventAccelerometerUpdate,
//}
//var gyroscopeEvents = []gods4.Event{
//	gods4.EventGyroscopeUpdate,
//}
//var batteryEvents = []gods4.Event{
//	// Battery
//	gods4.EventBatteryUpdate,
//}

// NewEvent returns a new Event and its associated data.
func NewGamepadEvent(name, action string, data interface{}) *GamepadEvent {
	return &GamepadEvent{Name: name, Action: action, Data: data}
}

type Gamepad interface {
	Projection() *GamepadProjection
	Start()
	Stop()
	EventChannel() *GamepadEventChannel
	//Rumble(rumble *Rumble)
	//Led(led *Led)
}

type GamepadProjection struct {
	Cross         bool
	Circle        bool
	Square        bool
	Triangle      bool
	L1            bool
	L2            byte
	L3            bool
	R1            bool
	R2            byte
	R3            bool
	DPadUp        bool
	DPadDown      bool
	DPadLeft      bool
	DPadRight     bool
	Share         bool
	Options       bool
	Ps            bool
	LeftStick     Stick
	RightStick    Stick
	Touchpad      Touchpad
	Accelerometer Accelerometer
	Gyroscope     Gyroscope
	Battery       Battery
	dPadDirection *float64
}

func NewGamepadProjection() *GamepadProjection {
	return &GamepadProjection{
		LeftStick:  Stick{},
		RightStick: Stick{},
		Touchpad: Touchpad{
			Press: false,
			Swipe: []Touch{
				{},
				{},
			},
		},
		Accelerometer: Accelerometer{},
		Gyroscope:     Gyroscope{},
		Battery:       Battery{},
	}
}

const RIGHT float64 = 0
const BOTTOM_RIGHT float64 = 45
const BOTTOM float64 = 90
const BOTTOM_LEFT float64 = 135
const LEFT float64 = 180
const TOP_LEFT float64 = 225
const TOP float64 = 270
const TOP_RIGHT float64 = 315

func (gp *GamepadProjection) updateDpadDirection() {
	var dPadDirection float64 = -1
	if gp.DPadUp && !gp.DPadDown {
		if gp.DPadRight && !gp.DPadLeft {
			dPadDirection = TOP_RIGHT
		} else if gp.DPadLeft && !gp.DPadRight {
			dPadDirection = TOP_LEFT
		} else {
			dPadDirection = TOP
		}
		gp.dPadDirection = &dPadDirection
	} else if gp.DPadDown && !gp.DPadUp {
		if gp.DPadRight && !gp.DPadLeft {
			dPadDirection = BOTTOM_RIGHT
		} else if gp.DPadLeft && !gp.DPadRight {
			dPadDirection = BOTTOM_LEFT
		} else {
			dPadDirection = BOTTOM
		}
		gp.dPadDirection = &dPadDirection
	} else if gp.DPadRight || gp.DPadLeft {
		if gp.DPadRight && !gp.DPadLeft {
			dPadDirection = RIGHT
		} else {
			dPadDirection = LEFT
		}
		gp.dPadDirection = &dPadDirection
	} else {
		gp.dPadDirection = nil
	}
}

func (gp *GamepadProjection) DPadDirection() *float64 {
	return gp.dPadDirection
}
