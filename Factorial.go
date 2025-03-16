package main

func Factorial(x int) int {
	if x > 0 {
		return Factorial(x-1) * x
	}
	return 1
}
