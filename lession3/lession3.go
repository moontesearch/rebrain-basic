package lession3

import "fmt"

var (
	A *int
	B int = 52
)

func Lession3() {
	A = &B
	fmt.Printf("%d", *A)
	*A = 5
	fmt.Printf("%d", B)
}
