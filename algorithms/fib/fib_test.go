package main

import "testing"

func TestFibOne(t *testing.T) {
	num := Fib(10)

	if num != 34 {
		t.Errorf("The fibonacci number was incorrect, got: %d, want: %d.", num, 34)
	}
}

func TestFibTwo(t *testing.T) {
	num := Fib(9)

	if num != 21 {
		t.Errorf("The fibonacci number was incorrect, got: %d, want: %d.", num, 21)
	}
}

func TestFibThree(t *testing.T) {
	num := Fib(-1)

	if num != 1 {
		t.Errorf("The fibonacci number was incorrect, got: %d, want: %d.", num, 1)
	}
}
