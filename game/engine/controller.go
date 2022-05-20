package engine

import (
	"github.com/gabz57/goledmatrix/controller"
	"time"
)

type (
	ControllerEngine struct {
		keyboard         controller.Keyboard
		keyboardEvents   []controller.KeyboardEvent
		gamepad          controller.Gamepad
		gamepadEvents    []controller.GamepadEvent
		activeController ControllerComponent
	}
	ControllerComponent interface {
		ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection)
		ConsumeKeyboardEvents(events *[]controller.KeyboardEvent, projection controller.KeyboardProjection)
		ProcessActions(*Engine) error
	}
	ActionFn func(e *Engine) error

	ControllerComponentBase struct {
		values  *EntityValues
		enabled bool
		actions []ActionFn
	}
)

func NewControllerEngine(keyboardChannel *controller.KeyboardEventChannel) *ControllerEngine {
	return &ControllerEngine{
		keyboard:         controller.NewKeyboardHard(keyboardChannel),
		gamepad:          controller.NewDualShock4(),
		keyboardEvents:   make([]controller.KeyboardEvent, controller.KeyboardEventChannelSize),
		gamepadEvents:    make([]controller.GamepadEvent, controller.GamepadEventChannelSize),
		activeController: nil,
	}
}

func (e *ControllerEngine) Start() {
	e.keyboard.Start()
	e.gamepad.Start()
}

func (e *ControllerEngine) Stop() {
	e.keyboard.Stop()
	e.gamepad.Stop()
}

func (e *ControllerEngine) SetActiveController(controller ControllerComponent) {
	e.activeController = controller
}

func (e *ControllerEngine) ConsumeKeyboardEvents(_ time.Duration) {
	// always renew consumed gamepadEvents
	e.keyboardEvents = e.keyboardEvents[:0]
eventReading:
	for {
		select {
		case event := <-(*e.keyboard.EventChannel()):
			e.keyboardEvents = append(e.keyboardEvents, *event)
		default:
			break eventReading
		}
	}
	// push event(s) to active controller
	if e.activeController != nil && len(e.keyboardEvents) > 0 {
		e.activeController.ConsumeKeyboardEvents(&e.keyboardEvents, *e.keyboard.Projection())
	}
}

func (e *ControllerEngine) ConsumeGamepadEvents(_ time.Duration) {
	// always renew consumed gamepadEvents
	e.gamepadEvents = e.gamepadEvents[:0]
eventReading:
	for {
		select {
		case event := <-(*e.gamepad.EventChannel()):
			e.gamepadEvents = append(e.gamepadEvents, *event)
		default:
			break eventReading
		}
	}
	// push event(s) to active controller
	if e.activeController != nil && len(e.gamepadEvents) > 0 {
		e.activeController.ConsumeGamepadEvents(&e.gamepadEvents, *e.gamepad.Projection())
	}
}

func (e *ControllerEngine) ProcessActions(engine *Engine, duration time.Duration) {
	if e.activeController != nil {
		err := e.activeController.ProcessActions(engine)
		if err != nil {
			println(err.Error())
			panic(err)
		}
	}
}

func NewControllerComponentBase(values *EntityValues) *ControllerComponentBase {
	return &ControllerComponentBase{
		values:  values,
		enabled: true,
		actions: nil,
	}
}

func (b *ControllerComponentBase) ClearActions() { b.actions = b.actions[:0] }
func (b *ControllerComponentBase) AddAction(action ActionFn) {
	b.actions = append(b.actions, action)
}
func (b *ControllerComponentBase) ProcessActions(e *Engine) error {
	if len(b.actions) > 0 {
		defer b.ClearActions()
		for _, actionFn := range b.actions {
			err := actionFn(e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *ControllerComponentBase) GetValue(ref ValueRef) interface{} {
	return b.values.Get(ref)
}

func (b *ControllerComponentBase) SetValue(ref ValueRef, value interface{}) {
	b.values.Set(ref, value)
}
