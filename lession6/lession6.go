package lession6

import "slices"

func Contains(a []string, x string) bool {
	return slices.Contains(a, x)
}

func GetMax() int {
	var a = make([]int, 0, 15)
	a = append(a, 1, 5, 6, 18, 1, 91, 1, 2, 4, 5, 6, 7, 8)
	return slices.Max(a)
}
