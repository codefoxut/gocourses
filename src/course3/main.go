package main

import (
	"fmt"
	// "course3/observer"
	"course3/state_machine"
)

func main() {
	// observer.ExecuteObserverPattern()
	// observer.ExecuteObserverPattern2()
	state_machine.ExecuteStateMachineClassic()
	// state_machine.ExecuteStateMachineHandmade()
	fmt.Println("Switch case machine.")
	state_machine.ExecuteStateMachineSwitch()
}
