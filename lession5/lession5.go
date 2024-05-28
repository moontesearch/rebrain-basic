package lession5

import (
	"fmt"
)

func MapRange() {
	var Books = map[string]map[string][]string{
		"HarryPotter": {
			"err": {
				"Oliver", "Mamammgag eagaae", "GOLANG NINJA"},
			"o": {"1", "III", "VI"},
		},
		"Harr": {
			"Okay": {
				"Hello World", "This lession 3", "Im so fain"},
			"OK": {"z", "g", "c"},
		},
	}
	fmt.Println(len(Books))

}
