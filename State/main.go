package main

import "fmt"

// light: on / off

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

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

// Every state has a standalone type
// Aggregate the BaseState.
type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

// Overwrite BaseState method to allow turnning Off
func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning the light off...")
	// replace the state with an OffState
	sw.State = NewOffState() // OnState switches itself into OffState
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

// Overwrite BaseState method to allow turnning On
func (o *OffState) On(sw *Switch) {
	fmt.Println("Turnning light on")
	sw.State = NewOnState() // OffState switches itself into OnState
}

func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off()
}
