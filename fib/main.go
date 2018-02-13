package main

import (
	"fmt"
)

var cache = make(map[int]int)

func Fib(n int) int {
	if n <= 3 {
		return 1
	}

	if _, ok := cache[n]; ok {
		return cache[n]
	}

	cache[n] = Fib(n-1) + Fib(n-2)

	return cache[n]
}

func main() {
	x := Fib(11)
	y := Fib(10)
	z := Fib(-1)
	d := Fib(50)

	fmt.Printf("x == %d\n", x)
	fmt.Printf("y == %d\n", y)
	fmt.Printf("z == %d\n", z)
	fmt.Printf("d == %d\n", d)
}
