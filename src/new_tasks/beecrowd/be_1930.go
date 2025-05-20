package beecrowd

func Be_1930(t [4]int) int {
	result := 0
	for i := 0; i < len(t); i++ {
		result += t[i]

	}
	return result - (len(t) - 1)

}
