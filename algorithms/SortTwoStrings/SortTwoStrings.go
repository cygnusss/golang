package main

import (
	"fmt"
	"strings"
)

func SortTwoStrings(str1, str2 string) string {
	out := []string{}
	l1, l2 := len([]rune(str1)), len([]rune(str2))
	letters := map[string]int{}

	for i := 0; i < l1; i++ {
		char := string(str1[i])
		if _, ok := letters[char]; ok {
			letters[char]++
		} else {
			letters[char] = 1
		}
	}

	for i := 0; i < l2; i++ {
		char := string(str2[i])
		if n, ok := letters[char]; ok {
			add := strings.Repeat(char, n)
			out = append(out, add)
			delete(letters, char)
		}
	}

	for i := 0; i < l1; i++ {
		char := string(str1[i])
		if _, ok := letters[char]; ok {
			out = append(out, char)
		}
	}

	return strings.Join(out, "")
}

func main() {
	fmt.Println(SortTwoStrings("program", "grapo"))
}
