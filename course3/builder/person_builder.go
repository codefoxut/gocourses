package person_builder

import (
	"fmt"
)

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	b.person.AnnualIncome = income
	return b
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Works().
		At("Goibibo.com").
		AsA("Software Engineer").
		Earning(10000).
		Lives().
		At("U block").
		In("Gurugram").
		WithPostcode("122002")

	person := pb.Build()
	fmt.Println(person)

}
