package main

import (
	"fmt"
	"math/big"

	"github.com/rargulati/algebra/fields"
)

// Exercise code paths. Remove when library is developed.
func main() {
	a, err := fields.NewFieldElement(big.NewInt(25), big.NewInt(57))
	if err != nil {
		panic(err)
	}
	b, err := fields.NewFieldElement(big.NewInt(19), big.NewInt(57))
	if err != nil {
		panic(err)
	}

	fmt.Printf("FieldElement a: %+v\n", a)
	fmt.Printf("FieldElement v: %+v\n", b)

	fmt.Printf("Does A(%d) equal B(%d)? %t\n", a.Value(), b.Value(), a.Equals(b))
	fmt.Printf("Does A(%d) equal A(%d)? %t\n", a.Value(), a.Value(), a.Equals(a))

	fmt.Printf("Does A(%d) not equal B(%d)? %t\n", a.Value(), b.Value(), a.NotEquals(b))
	fmt.Printf("Does A(%d) not equal A(%d)? %t\n", a.Value(), a.Value(), a.NotEquals(a))

	sumAB, err := a.Add(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("A(%d) added to B(%d): %d\n", a.Value(), b.Value(), sumAB.Value())

	sumAA, err := a.Add(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("A(%d) added to A(%d): %d\n", a.Value(), a.Value(), sumAA.Value())

	multAB, err := a.Mult(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("A(%d) multiplied with B(%d): %d\n", a.Value(), b.Value(), multAB.Value())

	multAA, err := a.Mult(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("A(%d) multiplied with A(%d): %d\n", a.Value(), a.Value(), multAA.Value())

	c, err := fields.NewFieldElement(big.NewInt(3), big.NewInt(31))
	if err != nil {
		panic(err)
	}
	d, err := fields.NewFieldElement(big.NewInt(24), big.NewInt(31))
	if err != nil {
		panic(err)
	}

	divCD, err := c.Div(d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("C(%d) divided by D(%d): %d\n", c.Value(), d.Value(), divCD.Value())

	c, err = fields.NewFieldElement(big.NewInt(17), big.NewInt(31))
	if err != nil {
		panic(err)
	}
	d, err = fields.NewFieldElement(big.NewInt(3), big.NewInt(31))
	if err != nil {
		panic(err)
	}

	divCD, err = c.Div(d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("C(%d) divided by D(%d): %d\n", c.Value(), d.Value(), divCD.Value())

	c, err = fields.NewFieldElement(big.NewInt(4), big.NewInt(31))
	if err != nil {
		panic(err)
	}
	d, err = fields.NewFieldElement(big.NewInt(4), big.NewInt(31))
	if err != nil {
		panic(err)
	}

	e, err := fields.NewFieldElement(big.NewInt(11), big.NewInt(31))
	if err != nil {
		panic(err)
	}

	divCD, err = c.Div(d)
	if err != nil {
		panic(err)
	}
	multCDE, err := divCD.Mult(e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("C(%d) divided by D(%d), then multiplied by E(%d): %d\n", c.Value(), d.Value(), e.Value(), multCDE.Value())
}
