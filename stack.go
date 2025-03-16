package main

import (
	"fmt"
	"math"
)

type Stack[T any] struct {
	MAX_SIZE int
	stack    []T
	ref      int
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.ref == -1
}

func (stack *Stack[T]) IsFull() bool {
	return stack.ref == stack.MAX_SIZE-1
}

func (stack *Stack[T]) Push(ele T) {
	if stack.IsFull() {
		fmt.Println("Stack is overflowed")
		return
	}
	stack.stack = append(stack.stack, ele)
	stack.ref++
}

func (stack *Stack[T]) Pop() interface{} {
	if stack.IsEmpty() {
		fmt.Println("stack Under flow")
		return math.MinInt
	}
	ele := stack.stack[stack.ref]
	stack.stack = stack.stack[:stack.ref]
	stack.ref--
	return ele
}

func DisplayStack[T any](items []T) {
	for _, val := range items {
		fmt.Printf("%v ", val)
	}
	fmt.Println()

}
