package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type CircularQueue struct {
	head *Node
	tail *Node
	size int
}

func (q *CircularQueue) Size() int {
	return q.size
}

func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue) Print() {
	current := q.tail.next
	count := 0
	for count < q.size {
		count++
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func (q *CircularQueue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Empty Queue")
		return 0
	}

	return q.tail.next.value
}

func (q *CircularQueue) Add(value int) {
	temp := &Node{value, nil}
	if q.tail == nil {
		q.tail = temp
		temp.next = q.tail
	} else {
		temp.next = q.tail.next
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *CircularQueue) Remove() int {
	if q.IsEmpty() {
		fmt.Println("Empty Queue")
		return 0
	}
	temp := q.tail.next
	q.tail.next = temp.next
	q.size--
	return temp.value
}

func main() {
	q := new(CircularQueue)
	q.Add(12)
	q.Add(15)
	q.Add(20)
	q.Remove()
	fmt.Println(q.Peek())
	q.Print()
}
