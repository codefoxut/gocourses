package factories

import (
	"fmt"
)

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, I am %s, I am %d years old!\n", p.name, p.age)
}

type rudePerson struct {
	name string
	age  int
}

func (p *rudePerson) SayHello() {
	fmt.Printf("Hi, I am not interested in talking to you.\n")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &rudePerson{name: name, age: age}
	}
	return &person{name: name, age: age}
}

func execute_factory() {
	p := NewPerson("James", 50)
	p.SayHello()

	p2 := NewPerson("Ruby", 101)
	p2.SayHello()

}
