package algebra2

func NumbersCase(n int) int {
	m := 0
	c := 0
	d := 0
	u := 0

	for n > 0 {
		if n > 1000 {
			n = n % 1000
			m++
		}
		if n < 1000 && n > 100 {
			n = n % 100
			c++
		}
		if n < 100 && n > 10 {
			n = n % 10
			d++
		} else {
			n -= 1
			u++
		}
	}

	return m

}

func Unic_Number(n int) int {

	return 1
}
