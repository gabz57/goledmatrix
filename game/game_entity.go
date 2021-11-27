package game

import "time"

const NoParent = -1

type (
	EntityRef int
	Entity    interface {
		Parent() EntityRef
		// Update consumes controller events & set directions/velocity
		Update(engine *Engine, dt time.Duration) bool
		// FinalUpdate apply computed positions to entities, if any
		FinalUpdate(dt time.Duration) bool
		Enable()
		Disable()
		GetController() ControllerComponent
	}

	ValueRef     int
	EntityValues map[ValueRef]interface{}

	EntityBase struct {
		Values    EntityValues
		ParentRef EntityRef
	}
)

func (ev EntityValues) Get(ref ValueRef) interface{} {
	return ev[ref]
}

func (ev EntityValues) Set(ref ValueRef, value interface{}) {
	ev[ref] = value
}

func NewEntityBase() *EntityBase {
	return &EntityBase{
		Values:    EntityValues{},
		ParentRef: NoParent,
	}
}

func (eb EntityBase) Parent() EntityRef {
	return eb.ParentRef
}

func (eb EntityBase) GetValue(ref ValueRef) interface{} {
	return eb.Values.Get(ref)
}

func (eb EntityBase) SetValue(ref ValueRef, value interface{}) {
	eb.Values.Set(ref, value)
}
