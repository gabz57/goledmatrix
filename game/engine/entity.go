package engine

import (
	"time"
)

const NoParent = -1

type (
	EntityRef int
	Entity    interface {
		Parent() EntityRef
		// Update consumes controller gamepadEvents & set directions/velocity
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

func (v EntityValues) Get(ref ValueRef) interface{} {
	return v[ref]
}

func (v EntityValues) Set(ref ValueRef, value interface{}) {
	v[ref] = value
}

func NewEntityBase() *EntityBase {
	return &EntityBase{
		Values:    EntityValues{},
		ParentRef: NoParent,
	}
}

func (b EntityBase) Parent() EntityRef {
	return b.ParentRef
}

func (b EntityBase) GetValue(ref ValueRef) interface{} {
	return b.Values.Get(ref)
}

func (b EntityBase) SetValue(ref ValueRef, value interface{}) {
	b.Values.Set(ref, value)
}
