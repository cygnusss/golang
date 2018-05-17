package main

// Tern does this
func Tern(cond bool, a, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
