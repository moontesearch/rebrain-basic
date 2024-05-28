package main

import (
	"fmt"

	"github.com/moontesearch/rebrain-basic/lession1"
	"github.com/moontesearch/rebrain-basic/lession2"
	"github.com/moontesearch/rebrain-basic/lession2/lession21"
	"github.com/moontesearch/rebrain-basic/lession3"
	"github.com/moontesearch/rebrain-basic/lession4"
	"github.com/moontesearch/rebrain-basic/lession5"
	"github.com/moontesearch/rebrain-basic/lession6"
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
	fmt.Println("Lession 4. slices package, slices.Clip(), slices.Delete()")
	lession4.SliceEdit()
	fmt.Println("Less 5, map[string]map[string][]string")
	lession5.MapRange()
	fmt.Println("Lession 6. contains and getMax")
	var a []string = []string{"Hello workd", "oaky gog", "y"}
	var s string = "y"
	fmt.Println("Contains", lession6.Contains(a, s))
	fmt.Println("getMax", lession6.GetMax())
}
