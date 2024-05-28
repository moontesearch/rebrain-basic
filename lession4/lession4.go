package lession4

import (
	"fmt"
)

func SliceEdit() {
	var (
		sliceWeek     = make([]string, 0, 7)
		sliceWeekWork = make([]string, 0, 5)
	)

	sliceWeek = append(sliceWeek,
		`Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`)
	fmt.Println("All items", sliceWeek)

	sliceWeekWork = sliceWeek[:5]
	fmt.Println("All items work", sliceWeekWork)

	sliceWeek = sliceWeek[5:]
	fmt.Println("Only weekends items", sliceWeek)

}
