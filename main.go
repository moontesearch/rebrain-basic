package main

import (
	"fmt"

	"github.com/moontesearch/rebrain-basic/lession1"
	"github.com/moontesearch/rebrain-basic/lession2"
	"github.com/moontesearch/rebrain-basic/lession2/lession21"
	"github.com/moontesearch/rebrain-basic/lession3"
)

func main() {
	fmt.Println("Lession 1. package time.Now()")
	lession1.FTime()
	fmt.Println("Lession 2. package strconv, strconv.Atoi(), strconv.Itoa()")
	lession2.AtoiItoa()
	fmt.Println("Lession 3. struct, interface, metod init interface")
	lession21.Lession21()
	fmt.Println("Lession 3.1 pointer, * and &")
	lession3.Radius()
	lession3.Lession3()
}
