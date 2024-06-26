package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)

		if math.Sqrt(x) == z {
			fmt.Println("OK")
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
