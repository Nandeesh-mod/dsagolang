package main

import "fmt"

func ResucrseFunc(x int) {
	if x > 0 {
		fmt.Printf("%v ", x)
		ResucrseFunc(x - 1)
		fmt.Println()
		fmt.Printf("%v ", x)
	}
}
