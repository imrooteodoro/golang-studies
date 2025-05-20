package numbers

import "fmt"

func dec_uni(n int) []int {
	arr := []int{}
	if n < 100 {
		dec := n / 10
		uni := n % 10
		arr = []int{dec, uni}
	} else {
		cent := n / 100
		dec := (n % 100) / 10
		uni := (n % 100) % 10
		arr = []int{cent, dec, uni}
	}

	return arr
}

func pow_func(arr []int) int {
	response := 0
	if len(arr) <= 2 {
		response = (arr[0] * arr[0]) + (arr[1] * arr[1])
	}
	if len(arr) > 2 {
		response = (arr[0] * arr[0]) + (arr[1] * arr[1]) + (arr[2] * arr[2])
	}
	return response
}

func Happy_number(n int) string {

	for n != 1 {
		res := dec_uni(n)
		n = pow_func(res)
		fmt.Printf("%v - %v\n", res, n)
	}

	return ""
}
