package main

import "fmt"

const MAX_SIZE = 5

type User struct {
	name string
	age  int
}

func main() {
	stack := Stack[User]{
		MAX_SIZE: MAX_SIZE,
		stack:    make([]User, 0, MAX_SIZE),
		ref:      -1,
	}

	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.IsFull())

	fmt.Println(stack.MAX_SIZE)

	stack.Push(User{
		name: "Nandeesh",
		age:  21,
	})

	stack.Push(User{
		name: "Mahesh",
		age:  23,
	})

	DisplayStack[User](stack.stack)
	stack.Pop()

	DisplayStack[User](stack.stack)

	ResucrseFunc(5)

	fmt.Printf("static val is : %v", StaticRecursion(5))
}
