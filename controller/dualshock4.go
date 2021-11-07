package controller

import (
	"fmt"
	"github.com/kpeu3i/gods4"
	"log"
	"strings"
	"time"
)

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

type DualShock4 struct {
	GamepadChannel    GamepadEventChannel
	projection        GamepadProjection
	controller        *gods4.Controller
	reconnectionTimer *time.Timer
	autoconnect       bool
}

func NewDualShock4() *DualShock4 {
	return &DualShock4{
		GamepadChannel: make(GamepadEventChannel, 1000),
		projection:     *NewGamepadProjection(),
	}
}

func (ds *DualShock4) EventChannel() *GamepadEventChannel {
	return &ds.GamepadChannel
}

func (ds *DualShock4) Start() {
	ds.autoConnect()
}

func (ds *DualShock4) Stop() {
	ds.autoconnect = false
	if ds.reconnectionTimer != nil {
		ds.reconnectionTimer.Stop()
	}
	if ds.controller != nil {
		log.Printf("* Controller #1 | %-10s | bye!\n", "Disconnect")
		_ = ds.controller.Disconnect()
	}
}

//
//func (ds *DualShock4) Rumble(rumble *Rumble) {
//	err := ds.controller.Rumble(rumble)
//	if err != nil {
//		return
//	}
//}
//
//func (ds *DualShock4) Led(led *Led) {
//	err := ds.controller.Led(led)
//	if err != nil {
//		return
//	}
//}

func (ds *DualShock4) autoConnect() {
	fmt.Println("autoConnect: Auto Connecting...")
	if !ds.performConnection() {
		ds.autoconnect = true
		go func() {
			fmt.Println("autoConnect: Start loop")
			for ds.autoconnect {
				fmt.Println("autoConnect: Will try in a moment")
				ds.reconnectionTimer = time.NewTimer(time.Second)
				select {
				case <-ds.reconnectionTimer.C:
					fmt.Println("autoConnect: Trying to connect")
					if ds.performConnection() {
						ds.autoconnect = false
					}
				}
			}
			fmt.Println("autoConnect: Stop loop")
		}()
	}
}

func (ds *DualShock4) performConnection() bool {
	if ds.connect(findController()) {
		ds.bindToGamepadChannel()
		go ds.listen()
		return true
	}
	return false
}

func (ds *DualShock4) listen() {
	fmt.Println("Start listening")
	err := ds.controller.Listen() // BLOCKING CALL
	fmt.Println("Stopped listening")

	if err != nil {
		println(err)
		go ds.autoConnect()
	}
}

func (ds *DualShock4) connect(controller *gods4.Controller) bool {
	if controller == nil {
		return false
	}
	err := controller.Connect()
	if err != nil {
		return false
	}
	ds.controller = controller

	log.Printf("* Controller #1 | %-10s | name: %s, connection: %s\n", "Connect", ds.controller, ds.controller.ConnectionType())
	//
	//// Disconnect controller when a program is terminated
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//go func() {
	//	<-signals
	//	err := ds.disconnect()
	//	if err != nil {
	//		//panic(err)
	//	}
	//	log.Printf("* Controller #1 | %-10s | bye!\n", "Disconnect")
	//}()
	log.Println("* Controller #1 CONNECTED")
	return true
}

// Find all controllers connected to your machine via USB or Bluetooth
func findController() *gods4.Controller {
	controllers := gods4.Find()
	if len(controllers) == 0 {
		return nil
	}
	// Select first controller from the list
	return controllers[0]
}

func (ds *DualShock4) Projection() *GamepadProjection {
	return &ds.projection
}

func (ds *DualShock4) bindToGamepadChannel() {
	println("bindToGamepadChannel")

	for i := range allEvents {
		event := allEvents[i]
		split := strings.Split(string(event), ".")
		ds.controller.On(event, func(data interface{}) error {
			switch event {
			case gods4.EventCrossPress:
				ds.projection.Cross = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Cross)
			case gods4.EventCrossRelease:
				ds.projection.Cross = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Cross)
			case gods4.EventCirclePress:
				ds.projection.Circle = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Circle)
			case gods4.EventCircleRelease:
				ds.projection.Circle = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Circle)
			case gods4.EventSquarePress:
				ds.projection.Square = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Square)
			case gods4.EventSquareRelease:
				ds.projection.Square = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Square)
			case gods4.EventTrianglePress:
				ds.projection.Triangle = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Triangle)
			case gods4.EventTriangleRelease:
				ds.projection.Triangle = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Triangle)
			case gods4.EventL1Press:
				ds.projection.L1 = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.L1)
			case gods4.EventL1Release:
				ds.projection.L1 = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.L1)
			case gods4.EventL2Press:
				ds.projection.L2 = data.(byte)
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.L2)
			case gods4.EventL2Release:
				ds.projection.L2 = 0
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.L2)
			case gods4.EventL3Press:
				ds.projection.L3 = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.L3)
			case gods4.EventL3Release:
				ds.projection.L3 = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.L3)
			case gods4.EventR1Press:
				ds.projection.R1 = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.R1)
			case gods4.EventR1Release:
				ds.projection.R1 = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.R1)
			case gods4.EventR2Press:
				ds.projection.R2 = data.(byte)
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.R2)
			case gods4.EventR2Release:
				ds.projection.R2 = 0
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.R2)
			case gods4.EventR3Press:
				ds.projection.R3 = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.R3)
			case gods4.EventR3Release:
				ds.projection.R3 = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.R3)
			case gods4.EventDPadUpPress:
				ds.projection.DPadUp = true
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadUp)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadUpRelease:
				ds.projection.DPadUp = false
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadUp)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadDownPress:
				ds.projection.DPadDown = true
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadDown)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadDownRelease:
				ds.projection.DPadDown = false
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadDown)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadLeftPress:
				ds.projection.DPadLeft = true
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadLeft)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadLeftRelease:
				ds.projection.DPadLeft = false
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadLeft)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadRightPress:
				ds.projection.DPadRight = true
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadRight)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventDPadRightRelease:
				ds.projection.DPadRight = false
				ds.projection.updateDpadDirection()
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.DPadRight)
				ds.GamepadChannel <- NewGamepadEvent("dpad", "direction", ds.projection.DPadDirection())
			case gods4.EventSharePress:
				ds.projection.Share = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Share)
			case gods4.EventShareRelease:
				ds.projection.Share = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Share)
			case gods4.EventOptionsPress:
				ds.projection.Options = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Options)
			case gods4.EventOptionsRelease:
				ds.projection.Options = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Options)
			case gods4.EventTouchpadSwipe:
				touchpad := data.(gods4.Touchpad)
				touch0 := touchpad.Swipe[0]
				touch1 := touchpad.Swipe[1]
				ds.projection.Touchpad.Swipe = []Touch{
					{
						IsActive: touch0.IsActive,
						X:        touch0.X,
						Y:        touch0.Y,
					},
					{
						IsActive: touch1.IsActive,
						X:        touch1.X,
						Y:        touch1.Y,
					},
				}
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Touchpad.Swipe)
			case gods4.EventTouchpadPress:
				ds.projection.Touchpad.Press = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Touchpad.Press)
			case gods4.EventTouchpadRelease:
				ds.projection.Touchpad.Press = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Touchpad.Press)
			case gods4.EventPSPress:
				ds.projection.Ps = true
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Ps)
			case gods4.EventPSRelease:
				ds.projection.Ps = false
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Ps)
			case gods4.EventLeftStickMove:
				stick := data.(gods4.Stick)
				ds.projection.LeftStick = Stick{X: stick.X, Y: stick.Y}
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.LeftStick)
			case gods4.EventRightStickMove:
				stick := data.(gods4.Stick)
				ds.projection.RightStick = Stick{X: stick.X, Y: stick.Y}
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.RightStick)
			case gods4.EventAccelerometerUpdate:
				acc := data.(gods4.Accelerometer)
				ds.projection.Accelerometer = Accelerometer{X: acc.X, Y: acc.Y, Z: acc.Z}
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Accelerometer)
			case gods4.EventGyroscopeUpdate:
				gyro := data.(gods4.Gyroscope)
				ds.projection.Gyroscope = Gyroscope{Roll: gyro.Roll, Yaw: gyro.Yaw, Pitch: gyro.Pitch}
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Gyroscope)
			case gods4.EventBatteryUpdate:
				battery := data.(gods4.Battery)
				ds.projection.Battery = Battery{
					Capacity:         battery.Capacity,
					IsCharging:       battery.IsCharging,
					IsCableConnected: battery.IsCableConnected,
				}
				ds.GamepadChannel <- NewGamepadEvent(split[0], split[1], ds.projection.Battery)
			}
			return nil
		})
	}
	println("bindToGamepadChannel DONE")
}
