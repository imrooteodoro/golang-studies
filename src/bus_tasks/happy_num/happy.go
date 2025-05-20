package happynum

import "fmt"

func DezUnid(n int) int {

	count_dez := 0

	count_uni := n % 10

	for n > 0 {

		n = n / 10
		count_dez++
	}

	fmt.Print(count_dez)

	return count_uni

}

func HappyNum(n int) string {

	v := DezUnid(n)

	//res := (v[0] * v[0])

	fmt.Print(v)

	//fmt.Print(res)

	return ""

}
