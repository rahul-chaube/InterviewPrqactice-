package stack

import "fmt"

type node struct {
	Data int
	node *node
}

type Stack struct {
	top *node
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(data int) {

	newNode := node{Data: data, node: s.top}
	s.top = &newNode

}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty ")

	}
	fmt.Println("Pop value is ", s.top.Data)
	s.top = s.top.node

}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty ")
		return -1
	}
	return s.top.Data
}

func (s *Stack) PrintAll() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty ")
	}

	curr := s.top

	for curr != nil {
		fmt.Println("Data is ", curr.Data)
		curr = curr.node
	}
}
