package state_machine

import (
	"os"
	"fmt"
	"strconv"
	"bufio"
)

type State2 int

const  (
	OffHook State2 = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s State2) String() string {
	switch s {
	case OffHook: return "OffHook"
	case Connecting: return "Connecting"
	case Connected: return "Connected"
	case OnHold: return "OnHold"
	case OnHook: return "OnHook"
	}
	return "Unknown"
  }

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed: return "CallDialed"
	case HungUp: return "HungUp"
	case CallConnected: return "CallConnected"
	case PlacedOnHold: return "PlacedOnHold"
	case TakenOffHold: return "TakenOffHold"
	case LeftMessage: return "LeftMessage"
	}
	return "Unknown"
  }

type TriggerResult struct {
	Trigger Trigger
	State State2
}


var rules  = map[State2][]TriggerResult {
	OffHook:  {
		{CallDialed, Connecting},
	},
	Connecting:  {
		{HungUp,  OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	  },
	  OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	  },
} 

func ExecuteStateMachineHandmade() {
	state, exitState := OffHook, OnHook
	for ok:= true; ok; ok = state !=  exitState {
		fmt.Println("the phone is currently", state)
		fmt.Println("Select a trigger:")

		for i := 0; i <len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), ".",  tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		j, _ := strconv.Atoi(string(input))

		tr :=  rules[state][j]
		state = tr.State
	}
	fmt.Println("we are done using the phone.")

}