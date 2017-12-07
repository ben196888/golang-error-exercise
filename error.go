package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0 ; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

type ErrNegativeSqrt float64

//func (e ErrNegativeSqrt) Error() string {
//	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
//}

func (e ErrNegativeSqrt) String() string {
	return fmt.Sprintf("To string negative number: %v", float64(e))
}

//func Sqrt(x float64) (float64, error) {
//	if x > 0.0 {
//		return sqrt(x), nil
//	}
//	return 0, ErrNegativeSqrt(x)
//}

func main() {
//	fmt.Println(Sqrt(2))
//	fmt.Println(Sqrt(-2))
	fmt.Println(ErrNegativeSqrt(4))
}

