package beecrowd

func Be_1153(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact

}
