package observer

import (
	"container/list"
	"fmt"
)

type PropertyChange struct {
	Name  string
	Value interface{}
}

type PersonAge struct {
	Observable
	age int
}

func NewPersonAge(age int) *PersonAge {
	return &PersonAge{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *PersonAge) Age() int {
	return p.age
}

func (p *PersonAge) SetAge(age int) {
	if age == p.age {
		return
	}
	fmt.Println("setting the age to", age)

	oldCanVote := p.CanVote()
	p.age = age
	p.Fire(PropertyChange{"age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *PersonAge) CanVote() bool {
	return p.age >= 18
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 16 {
			fmt.Println("Congratulations, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

type ElectoralRoll struct{}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func ExecuteObserverPattern2() {
	// traffic management
	p2 := NewPersonAge(15)
	ts := &TrafficManagement{p2.Observable}
	er := &ElectoralRoll{}
	p2.Subscribe(ts)
	p2.Subscribe(er)
	for i := 16; i <= 20; i++ {
		p2.SetAge(i)
	}

	//

}
