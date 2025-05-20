package algebra2

import (
	"sort"
)

func MajorProduct(arr []int) int {
	sort.Ints(arr)

	pos1 := len(arr) - 1
	pos2 := len(arr) - 2

	result := arr[pos1] * arr[pos2]

	return result
}
