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

	fmt.Println(a.Equals(b))
	fmt.Println(a.Equals(a))

	fmt.Println(a.NotEquals(b))
	fmt.Println(a.NotEquals(a))

	fmt.Println(a)
	fmt.Println(b)
}
