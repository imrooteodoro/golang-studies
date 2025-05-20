package letters

import (
	"fmt"
	"strings"
)

func Semantic(word string) []int {
	vog := []string{"a", "e", "i", "o", "u"}
	word_array := strings.Split(word, "")

	vog_count := 0
	cons_count := 0

	for i := range word_array {
		if strings.Contains(word, vog[i]) {
			fmt.Print(vog[i])
			vog_count += 1
		} else {
			cons_count += 1
		}

	}
	arr := []int{}
	return append(arr, cons_count, vog_count)

}
