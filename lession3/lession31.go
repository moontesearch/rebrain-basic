package lession3

import (
	"fmt"
	"math"
)

func Radius() {
	var r float64 = 35 / (2 * 3.14)
	var rPtr *float64 = &r
	fmt.Printf("Radius: %.2f\n", *rPtr)
	var s float64 = 3.14 * math.Pow(*rPtr, 2)
	fmt.Printf("Circle S = %.2f\n", s)
}
