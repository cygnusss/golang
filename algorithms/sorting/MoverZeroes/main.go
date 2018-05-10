package main

import (
	"fmt"
)

func MoveZeroes(nums []int) []int {
	left := 0              // 1
	right := len(nums) - 1 // 1

	for right > left { // n - 1 + 1
		if nums[right] == 0 {
			right--
			continue
		}
		if nums[left] == 0 {
			nums[left] = nums[right]
			nums[right] = 0
		}
		left++
	}

	return nums
}

func main() {
	t1 := MoveZeroes([]int{0, 1, 1, 2})
	t2 := MoveZeroes([]int{1, 2, 0, 0})
	t3 := MoveZeroes([]int{0, 2})
	t4 := MoveZeroes([]int{1, 1, 0, 1, 0})

	fmt.Println("Test #1:", t1)
	fmt.Println("Test #2:", t2)
	fmt.Println("Test #3:", t3)
	fmt.Println("Test #4:", t4)
}
