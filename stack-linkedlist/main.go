package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

func (stack *StackLinkedList) isEmpty() bool {
	return stack.size == 0
}

func (stack *StackLinkedList) push(value int) {
	stack.head = &Node{value, stack.head}
	stack.size++
}

func (stack *StackLinkedList) peek() int {
	if stack.isEmpty() {
		return 0
	}

	return stack.head.value
}

func (stack *StackLinkedList) pop() (int, bool) {
	if stack.isEmpty() {
		return 0, false
	}
	value := stack.head.value
	stack.head = stack.head.next
	stack.size--

	return value, true
}

func (stack *StackLinkedList) print() {
	temp := stack.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}

	fmt.Println()
}

func main() {
	stack := new(StackLinkedList)
	stack.push(12)
	stack.push(20)
	stack.print()
	fmt.Println(stack.peek())
}
