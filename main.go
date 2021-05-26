package goreverser

import (
	"strings"
)

func Reverse(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	println("String reverser")

	s := "1 and 2 and 3"
	reverseString := Reverse(s)

	println(reverseString)
}
