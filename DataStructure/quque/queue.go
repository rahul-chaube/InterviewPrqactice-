package quque

import "fmt"

type node struct {
	data int
	node *node
}

type Queue struct {
	tail *node
	head *node
}

func (q *Queue) IsEmpty() bool {

	return q.head == nil

}

func (q *Queue) Enqueue(data int) {

	newNode := node{data: data}

	if q.tail != nil {
		q.tail.node = &newNode
	}

	q.tail = &newNode

	if q.head == nil {
		q.head = &newNode
	}

}

func (q *Queue) Dequeue() {

	if q.IsEmpty() {
		fmt.Println("Queue is empty ")
	}

	fmt.Println("Dequue value is ", q.head.data)

	q.head = q.head.node

	if q.head == nil {
		q.tail = nil
	}

}

func (q *Queue) PrintAll() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty ")
	}

	curr := q.head

	for curr != nil {
		fmt.Println("Data is ", curr.data)
		curr = curr.node
	}
}
