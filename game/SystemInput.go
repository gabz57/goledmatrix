package game

import (
	"time"
)

type (
	InputSystem struct {
		inputs []InputComponent
	}

	InputComponent interface {
		GetEntity() Entity
		UpdateInputs(elapsedBetweenUpdate time.Duration)
	}
)

func NewInputSystem() *InputSystem {
	return &InputSystem{}
}

func (s InputSystem) Add(input *InputComponent) {
	if input != nil {
		s.inputs = append(s.inputs, *input)
	}
}

func (s InputSystem) Update(w *World) {
	// check each InputComponent for new user input to process

	// deduce Action from input

	// transform Action to Command (depends on context)

	// then apply Command to the "system"

	//for _, inputComponent := range s.inputSystem {
	//
	//}
}
