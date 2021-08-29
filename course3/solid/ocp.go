package solid

// ocp
// open for extension, closed for modification
// Specification

import (
	"fmt"
)

type Color int
type Size int

const (
	red Color = iota
	green
	blue
	yellow
)

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// this should not be done this way, as we need to update the filtering funtions or add new filter functions which are basically the samee functions.

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type ColorAndSizeSpecification struct {
	color Color
	size  Size
}

func (s *ColorAndSizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size && p.color == s.color
}

type AndSpecification struct {
	first, second Specification
}

func (s *AndSpecification) IsSatisfied(p *Product) bool {
	return s.first.IsSatisfied(p) && s.second.IsSatisfied(p)
}

type ProductFilter struct{}

func (pf *ProductFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	for _, v := range f.filterBySize(products, small) {
		fmt.Printf(" - %s is small\n", v.name)
	}

	for _, v := range f.filterBySizeAndColor(products, small, green) {
		fmt.Printf(" - %s is small and green\n", v.name)
	}

	fmt.Println("New implementation.")
	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	gsSpec := ColorAndSizeSpecification{green, small}
	glSpec := AndSpecification{&greenSpec, &largeSpec}

	pf1 := ProductFilter{}
	for _, v := range pf1.Filter(products, &greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	for _, v := range pf1.Filter(products, &largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}
	for _, v := range pf1.Filter(products, &gsSpec) {
		fmt.Printf(" - %s is small and green\n", v.name)
	}

	for _, v := range pf1.Filter(products, &glSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

}
