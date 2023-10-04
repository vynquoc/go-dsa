package main

import "fmt"

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]interface{}
	front int
	back  int
}

func (q *Queue) enqueue(value interface{}) {
	if q.size >= capacity {
		return
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
}

func (q *Queue) dequeue() interface{} {
	var value interface{}
	if q.size <= 0 {
		return 0
	}
	q.size--
	value = q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value
}

func (q *Queue) length() int {
	return q.size
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func main() {
	q := new(Queue)
	q.enqueue(7)
	q.enqueue(12)
	for q.isEmpty() == false {
		val, _ := q.dequeue().(int)
		fmt.Println(val)
	}
}
