package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("ошибка: отриц. значение %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 1; ; i++ {
		lastZ := z
		z -= (z*z - x) / (2 * z)

		fmt.Println(i, " шаг: ", z)

		if math.Abs(lastZ-z) < 0.00001 {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
