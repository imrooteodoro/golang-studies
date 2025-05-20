package recursive

func Fac(n int) int {

	if n >= 0 && n <= 1 {
		return 1
	}

	return n * Fac(n-1)
}

//5 x 4 = 20
//20 x 3 = 60
//60 x 2 = 120
