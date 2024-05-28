package lession2

import (
	"fmt"
	"strconv"
)

func AtoiItoa() {
	var (
		a   string = "104"
		aa  int
		b   int = 35
		bb  string
		err error
	)

	aa, err = strconv.Atoi(a)
	if err != nil {
		panic("a != int")
	}
	fmt.Printf("%T", aa)

	strconv.Itoa(b)
	fmt.Printf("%T", bb)

}
