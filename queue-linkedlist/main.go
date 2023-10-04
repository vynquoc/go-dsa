package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type QueueLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Empty Queue")
		return 0
	}
	return q.head.value
}

func (q *QueueLinkedList) Print() {
	current := q.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func (q *QueueLinkedList) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) Remove() int {
	if q.IsEmpty() {
		fmt.Println("Empty Queue")
		return 0
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value
}

func main() {
	q := new(QueueLinkedList)
	q.Add(4)
	q.Add(5)
	q.Remove()
	q.Print()
}
