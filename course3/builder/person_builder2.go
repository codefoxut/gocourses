package person_builder2

import (
	"fmt"
)

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, action := range b.actions {
		action(&p)
	}
	return &p
}

func main() {

	pb := PersonBuilder{}
	p := pb.
		Called("Ujjwal").
		WorksAsA("Software Engineer").
		Build()

	fmt.Println(*p)

}
