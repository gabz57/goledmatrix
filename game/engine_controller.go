package game

import (
	"github.com/gabz57/goledmatrix/controller"
	"time"
)

type (
	ControllerEngine struct {
		gamepad          controller.Gamepad
		events           []controller.GamepadEvent
		activeController ControllerComponent
	}
	ControllerComponent interface {
		ConsumeGamepadEvents(events *[]controller.GamepadEvent, projection controller.GamepadProjection)
		ProcessActions(*Engine) error
	}
	ActionFn func(e *Engine) error

	ControllerComponentBase struct {
		values  *EntityValues
		enabled bool
		actions []ActionFn
	}
)

func NewControllerEngine() *ControllerEngine {
	return &ControllerEngine{
		gamepad:          controller.NewDualShock4(),
		events:           make([]controller.GamepadEvent, controller.GamepadEventChannelSize),
		activeController: nil,
	}
}

func (ce *ControllerEngine) Start() {
	ce.gamepad.Start()
}

func (ce *ControllerEngine) Stop() {
	ce.gamepad.Stop()
}

func (ce *ControllerEngine) SetActiveController(controller ControllerComponent) {
	ce.activeController = controller
}

func (ce *ControllerEngine) ConsumeGamepadEvents(_ time.Duration) {
	// always renew consumed events
	ce.events = ce.events[:0]
eventReading:
	for {
		select {
		case event := <-(*ce.gamepad.EventChannel()):
			ce.events = append(ce.events, *event)
		default:
			break eventReading
		}
	}
	// push event(s) to active controller
	if ce.activeController != nil && len(ce.events) > 0 {
		projection := *ce.gamepad.Projection()
		ce.activeController.ConsumeGamepadEvents(&ce.events, projection)
	}
}

func (ce *ControllerEngine) ProcessActions(e *Engine, duration time.Duration) {
	if ce.activeController != nil {
		err := ce.activeController.ProcessActions(e)
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

func (ccb *ControllerComponentBase) ClearActions() { ccb.actions = ccb.actions[:0] }
func (ccb *ControllerComponentBase) AddAction(action ActionFn) {
	ccb.actions = append(ccb.actions, action)
}
func (ccb *ControllerComponentBase) ProcessActions(e *Engine) error {
	if len(ccb.actions) > 0 {
		defer ccb.ClearActions()
		for _, actionFn := range ccb.actions {
			err := actionFn(e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (ccb *ControllerComponentBase) GetValue(ref ValueRef) interface{} {
	return ccb.values.Get(ref)
}

func (ccb *ControllerComponentBase) SetValue(ref ValueRef, value interface{}) {
	ccb.values.Set(ref, value)
}
