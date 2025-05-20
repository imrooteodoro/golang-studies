package letters

import (
	"slices"
	"strings"
)

func Palindrones(word string) string {

	pali := strings.Split(word, "")
	word2 := strings.Split(word, "")

	slices.Reverse(pali)

	if word2[0] == pali[0] {
		return "SIM\n"
	}

	return "NAO\n"

}
