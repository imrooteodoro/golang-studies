package letters

import (
	"strings"
)

func CommPrefix(word [3]string) string {

	count := 0

	for i := range word {
		strings.ToLower(word[i])
	}

	for i := 0; i < len(word[0]); i++ {
		if i < len(word[1]) && i < len(word[2]) && word[0][0:i+1] == word[1][0:i+1] && word[0][0:i+1] == word[2][0:i+1] {
			count += 1
		}
	}
	//fmt.Printf("%v\n", count)
	return word[0][0:count]
}
