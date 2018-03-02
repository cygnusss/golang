package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Is Palindrome 5225: %v\n", isPalindrome(5225))
	fmt.Printf("Is Palindrome 1331: %v\n", isPalindrome(1331))
	fmt.Printf("Is Palindrome 2333: %v\n", isPalindrome(2333))
}

func isPalindrome(n int) bool {
	fin := 0
	tmp := n

	for n > 0 {
		l := n % 10
		fin = fin*10 + l
		n /= 10
	}

	if tmp == fin {
		return true
	}

	return false
}
