package controller

const RIGHT float64 = 0
const BOTTOM_RIGHT float64 = 45
const BOTTOM float64 = 90
const BOTTOM_LEFT float64 = 135
const LEFT float64 = 180
const TOP_LEFT float64 = 225
const TOP float64 = 270
const TOP_RIGHT float64 = 315

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

func (gp *GamepadProjection) updateDpadDirection() {
	dPadDirection, none := direction(gp.DPadUp, gp.DPadDown, gp.DPadRight, gp.DPadLeft)
	if none {
		gp.dPadDirection = nil
	} else {
		gp.dPadDirection = &dPadDirection
	}
}

func (gp *GamepadProjection) DPadDirection() *float64 {
	return gp.dPadDirection
}

func direction(up bool, down bool, right bool, left bool) (float64, bool) {
	if up && !down {
		if right && !left {
			return TOP_RIGHT, false
		} else if left && !right {
			return TOP_LEFT, false
		} else {
			return TOP, false
		}
	} else if down && !up {
		if right && !left {
			return BOTTOM_RIGHT, false
		} else if left && !right {
			return BOTTOM_LEFT, false
		} else {
			return BOTTOM, false
		}
	} else if right || left {
		if right && !left {
			return RIGHT, false
		} else {
			return LEFT, false
		}
	} else {
		// no direction
		return -1, true
	}
}
