package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt func 's parameter should be non-negative, instread you provide %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	for math.Abs(x-z*z) >= 1e-10 {
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	val, err := Sqrt(-2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}
