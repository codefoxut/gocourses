package main

// dependency Inversion Principle
// High level module not depend on Low level module
// both should depend on abstraction

import (
	"fmt"
)

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// ...
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for _, v := range rs.relations {
		if v.from.name == name && v.relationship == Parent {
			result = append(result, v.to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

type BadResearch struct {
	relationships Relationships
}

func (br *BadResearch) Investigate() {
	relations := br.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}

type Research struct {
	browser RelationshipBrowser
}

func (br *Research) Investigate() {
	for _, child := range br.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", child.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	// low-level module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	badResearch := BadResearch{relationships}
	badResearch.Investigate()

	fmt.Println("good research.")
	research := Research{&relationships}
	research.Investigate()
}
