package main

import (
	"fmt"
)

func FromRoman(num string) int {
	romans := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	n := 0

	for i, val := range num {
		next := 0

		cur := romans[string(val)]
		if i != len(num)-1 {
			next = romans[string(num[i+1])]
		}

		if cur < next {
			n -= cur
		} else {
			n += cur
		}
	}

	return n
}

func main() {
	t1 := FromRoman("X")
	t2 := FromRoman("IX")
	t3 := FromRoman("MXII")
	t4 := FromRoman("CMXLV")

	fmt.Printf("Test 1: %d\n", t1)
	fmt.Printf("Test 2: %d\n", t2)
	fmt.Printf("Test 3: %d\n", t3)
	fmt.Printf("Test 4: %d\n", t4)
}
