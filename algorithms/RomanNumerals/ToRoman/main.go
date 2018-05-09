package main

import (
	"fmt"
	"sort"
	"strings"
)

func ToRoman(num int) string {
	// Necessary to keep track of numers as
	// CM, CD, XC, XL, IX, IV
	// Because except for these parts
	// logic for roman nums is the same as
	// regular base 5 numbers
	romanmap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	var romans []int

	for k := range romanmap {
		romans = append(romans, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(romans)))

	result := ""
	// num 1000
	for _, ar := range romans {
		q := num / ar  // num 1000/1000 = 1
		num = num % ar // num 1900%1100 = 0

		rm := romanmap[ar]

		result += strings.Repeat(rm, q)
	}

	return result
}

func main() {
	t1 := ToRoman(5)
	t2 := ToRoman(12)
	t3 := ToRoman(1300)
	t4 := ToRoman(999)

	fmt.Printf("Test 1: %s\n", t1)
	fmt.Printf("Test 2: %s\n", t2)
	fmt.Printf("Test 3: %s\n", t3)
	fmt.Printf("Test 4: %s\n", t4)
}
