package main

import "fmt"

type Stack struct {
	s []int
}

func (stack *Stack) isEmpty() bool {
	length := len(stack.s)
	return length == 0
}

func (stack *Stack) length() int {
	length := len(stack.s)
	return length
}

func (stack *Stack) print() {
	length := stack.length()
	for i := 0; i < length; i++ {
		fmt.Print(stack.s[i], " ")
	}

	fmt.Println()
}

func (stack *Stack) push(value int) {
	stack.s = append(stack.s, value)
}

func (stack *Stack) pop() int {
	if stack.isEmpty() {
		return 0
	}

	length := stack.length()
	res := stack.s[length-1]
	stack.s = stack.s[:length-1]
	return res
}

func (stack *Stack) top() int {
	if stack.isEmpty() {
		return 0
	}
	length := stack.length()
	return stack.s[length-1]
}

func main() {
	stack := new(Stack)
	stack.push(25)
	stack.push(40)
	fmt.Println(stack.top())
	stack.print()
}
