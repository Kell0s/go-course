package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 1; ; i++ {
		lastZ := z
		z -= (z*z - x) / (2 * z)

		fmt.Println(i, " шаг: ", z)

		if math.Abs(lastZ-z) < 0.00001 {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(9))
}
