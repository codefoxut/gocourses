package main

import (
	// "course3/observer"
	// "course3/state_machine"
	// "course3/strategy"
	// "course3/template_pattern"
	"course3/visitor"
)

func main() {
	// observer.ExecuteObserverPattern()
	// observer.ExecuteObserverPattern2()
	// state_machine.ExecuteStateMachineClassic()
	// state_machine.ExecuteStateMachineHandmade()
	// fmt.Println("Switch case machine.")
	// state_machine.ExecuteStateMachineSwitch()
	// strategy.ExecuteStrategyPattern()
	// template_pattern.ExecuteTemplateMethod()
	// template_pattern.ExecuteTemplateFunction()
	visitor.ExecuteVisitorIntrusive()
	visitor.ExecuteVisitorReflective()
	visitor.ExecuteVisitorClassic()
}
