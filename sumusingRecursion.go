package main

func SumRecursion(x int) int {
	if x > 0 {
		return SumRecursion(x-1) + x
	}
	return 0
}
