package numbers

func digits(n int) []int {
	arr_response := []int{}

	mil := n / 1000
	cent := (n % 1000) / 100
	dec := ((n % 1000) % 100) / 10
	uni := ((n % 1000) % 100) % 10
	arr_response = append(arr_response, mil, cent, dec, uni)
	return arr_response

}

func Sum_digits(n int) int {
	result := 0
	res := digits(n)
	arr_res := len(res)

	for arr_res != 1 {
		result = res[0] + res[1] + res[2] + res[3]
		res = digits(result)
		if res[0] == 0 && res[1] == 0 && res[2] == 0 {
			arr_res = 1
		}

	}

	return result

}
