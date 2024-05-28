package main

import (
	"fmt"

	"github.com/moontesearch/rebrain-basic/lession1"
	"github.com/moontesearch/rebrain-basic/lession2"
	"github.com/moontesearch/rebrain-basic/lession2/lession21"
)

func main() {
	fmt.Println("Lession 1. package time.Now()")
	lession1.FTime()
	fmt.Println("Lession 2. package strconv, strconv.Atoi(), strconv.Itoa()")
	lession2.AtoiItoa()
	fmt.Println("Lession 3. struct, interface, metod init interface")
	lession21.Lession21()
}
