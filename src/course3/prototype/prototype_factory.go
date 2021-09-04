package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}


func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	// fmt.Println(string(b.Bytes()))

	q := Person{}
	d := gob.NewDecoder(&b)
	_ = d.Decode(&q)

	return &q

}


type Office struct {
	Suite int
	StreetAddress, City string
  }
  
type Employee struct {
	Name string
	Office Office
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	// fmt.Println(string(b.Bytes()))

	q := Employee{}
	d := gob.NewDecoder(&b)
	_ = d.Decode(&q)

	return &q

}

// employee factory
var mainOffice = Employee{"", Office{0, "123 East Dr", "London"}}
var auxOffice = Employee{"", Office{0, "66 West Dr", "London"}}

func newEmployee(proto *Employee, name string, suite int) *Employee{
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	  return newEmployee(&mainOffice, name, suite)
  }
  
func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func execute() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt", "Sam"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Jill")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	johny := NewMainOfficeEmployee("Johny", 100)
  	janet := NewAuxOfficeEmployee("Janet", 200)

  	fmt.Println(johny)
  	fmt.Println(janet)

}
