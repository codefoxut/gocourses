package state_machine

import (
	"fmt"
)

type Switch struct {
	State State
}

func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct {}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("light is already on.")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("light is already off.")
}

type OnState struct {
	BaseState
}

func NewOnState()  *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (b *OnState) Off(sw *Switch) {
	fmt.Println("turning off light.")
	sw.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState()  *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (b *OffState) On(sw *Switch) {
	fmt.Println("turning on light.")
	sw.State = NewOnState()
}


func ExecuteStateMachineClassic() {
	sw := NewSwitch()
  sw.On()
  sw.Off()
  sw.Off()
  sw.On()
}