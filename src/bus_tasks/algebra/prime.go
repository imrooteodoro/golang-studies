package algebra

func Prime(n int) string {

	if n == 2 {
		return "PRIMO"
	} else {
		for i := 3; i < n; i++ {
			if n%i == 0 {
				return "NAO EH PRIMO"
			}
		}
	}

	return "PRIMO"
}
