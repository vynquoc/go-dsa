package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Length() int {
	return len(s.data)
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	length := len(s.data)
	res := s.data[length-1]
	s.data = s.data[:length-1]
	return res
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return 0
	}
	return s.data[len(s.data)-1]
}

// QUEUE
type StackQueue struct {
	stk1 Stack
	stk2 Stack
}

func (q *StackQueue) Add(value int) {
	q.stk1.Push(value)
}

func (q *StackQueue) Remove() int {
	var value int
	if q.stk2.IsEmpty() == false {
		value = q.stk2.Pop()
		return value
	}
	for q.stk1.IsEmpty() == false {
		q.stk2.Push(q.stk1.Pop())
	}
	value = q.stk2.Pop()
	return value
}

func (q *StackQueue) Length() int {
	return q.stk1.Length() + q.stk2.Length()
}

func main() {
	q := new(StackQueue)
	q.Add(1)
	q.Add(2)
	q.Add(3)
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	q.Add(6)
	fmt.Println(q)
}
