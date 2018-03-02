package main

import (
	"fmt"
)

func main() {
	houses := []int{8, 9, 10}
	res := rob(houses)

	fmt.Printf("The best option is: %d\n", res)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		return nums[0]
	}

	a := 0
	b := 0

	for i := 0; i < len(nums); i++ {
		c := b + nums[i]
		if c < a {
			c = a
		}

		b = a
		a = c
	}

	return a
}
