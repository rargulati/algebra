package fields

import (
	"fmt"
	"math/big"
)

// Element in a Finite Field of order _prime_
type FieldElement struct {
	// specific element in the field, contained in Fprime
	num *big.Int
	// order of the field
	prime *big.Int
}

func NewFieldElement(num *big.Int, prime *big.Int) (*FieldElement, error) {
	// If num is less than 0 OR num is greater than or equal to prime, we error
	if num.Cmp(big.NewInt(0)) == -1 || (num.Cmp(prime) == 0 || num.Cmp(prime) == 1) {
		return nil, fmt.Errorf(
			"num should be between 0 and %d - 1; got %d",
			prime.Int64(),
			num.Int64())
	}

	return &FieldElement{num: num, prime: prime}, nil
}

func (fe *FieldElement) Value() int64 {
	return fe.num.Int64()
}

func (fe *FieldElement) Order() int64 {
	return fe.prime.Int64()
}

// Equals tests for equality between two FieldElements
func (fe *FieldElement) Equals(fe2 *FieldElement) bool {
	if fe.num.Cmp(fe2.num) == 0 && fe.prime.Cmp(fe2.prime) == 0 {
		return true
	}
	return false
}

// NotEquals tests for inequality between two FieldElements
func (fe *FieldElement) NotEquals(fe2 *FieldElement) bool {
	if fe.num.Cmp(fe2.num) != 0 || fe.prime.Cmp(fe2.prime) != 0 {
		return true
	}
	return false
}

// Override String() for FieldElement
func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement num: %d, prime: %d",
		fe.num.Int64(),
		fe.prime.Int64())
}

// Add holds true if and only if for FieldElements A and B, A and B are elements
// of Fprime, and then A.Add(B) is an element of Fprime. Field addition must be
// be commutative and associative, with the values being closed under the order.
func (fe *FieldElement) Add(fe2 *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(fe2.prime) != 0 {
		return nil, fmt.Errorf(
			"can't add two numbers of different Fields; expected order %d",
			fe.prime.Int64())
	}

	sum := big.NewInt(0).Add(fe.num, fe2.num)
	mod := big.NewInt(0).Mod(sum, fe.prime)

	return &FieldElement{num: mod, prime: fe.prime}, nil
}

// Sub holds true if and only if for FieldElements A and B, A and B are elements
// of Fprime, and then A.Sub(B) is an element of Fprime. Field subtraction must
// be commutative and associative, with the values being closed under the order.
func (fe *FieldElement) Sub(fe2 *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(fe2.prime) != 0 {
		return nil, fmt.Errorf(
			"can't subtract two numbers of different Fields; expected order %d",
			fe.prime.Int64())
	}

	sub := big.NewInt(0).Sub(fe.num, fe2.num)
	mod := big.NewInt(0).Mod(sub, fe.prime)

	return &FieldElement{num: mod, prime: fe.prime}, nil
}

// Mult holds true if and only if for FieldElements A and B, A and B are elements
// of Fprime, and then A.Mult(B) is an element of Fprime. Field multiplication must
// be commutative and associative, with the values being closed under the order.
// Mult between A and B is akin to executing field addition of A, B times.
func (fe *FieldElement) Mult(fe2 *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(fe2.prime) != 0 {
		return nil, fmt.Errorf(
			"can't multiply two numbers of different Fields; expected order %d",
			fe.prime.Int64())
	}

	mult := big.NewInt(0).Mul(fe.num, fe2.num)
	mod := big.NewInt(0).Mod(mult, fe.prime)

	return &FieldElement{num: mod, prime: fe.prime}, nil
}

// Exp of a FieldElement A takes an int B, which is the exponent, and executes
// field multiplication of A, B times. This value is then modulus the order of A,
// Fprime, as the product of A^B is bound to the order of A.
func (fe *FieldElement) Exp(exponent *big.Int) (*FieldElement, error) {
	primeMinusOne := big.NewInt(0).Sub(fe.prime, big.NewInt(1))
	// adjust for negative exponents by exploiting a^(p-1) = 1
	fieldValue := big.NewInt(0).Mod(exponent, primeMinusOne)
	// big int exponentiation takes the modulus into consideration
	exp := big.NewInt(0).Exp(fe.num, fieldValue, fe.prime)
	return &FieldElement{num: exp, prime: fe.prime}, nil
}

func (fe *FieldElement) Div(fe2 *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(fe2.prime) != 0 {
		return nil, fmt.Errorf(
			"can't divide two numbers of different Fields; expected order %d",
			fe.prime.Int64())
	}

	// employ exponentiation via fermats little theorem
	exponent := big.NewInt(0).Sub(fe.prime, big.NewInt(2))
	exp := big.NewInt(0).Exp(fe2.num, exponent, big.NewInt(0))
	// produce the product by multiplying against the denominator
	product, err := fe.Mult(&FieldElement{num: exp, prime: fe.prime})
	if err != nil {
		return nil, err
	}
	return &FieldElement{num: product.num, prime: product.prime}, nil
}
