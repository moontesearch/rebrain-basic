package lession2

import (
	"fmt"
	"strconv"
)

var (
	A = "104"
	B = 35
)

func AtoiItoa() {

	a, _ := strconv.Atoi(A)
	fmt.Printf("strconv.Atoi() - string in int: %T\n", a)

	b := strconv.Itoa(B)
	fmt.Printf("strconv.Atoi() - int in string: %T\n", b)

}
