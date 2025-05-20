package arrays

import (
	"sort"
)

func MajorProduct(arr []int) int {

	sort.Slice(arr, func(i, j int) bool {
		arr := arr[i] > arr[j]
		return arr
	})

	//fmt.Print(arr)
	return arr[0] * arr[1]

}
