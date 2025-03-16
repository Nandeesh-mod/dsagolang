package main

var StaticVal int = 0

func StaticRecursion(x int) int {
	if x > 0 {
		StaticVal += 1
		return StaticRecursion(x-1) + StaticVal
	}
	return 0
}
