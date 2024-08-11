package doublelinkedlist

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type Doubly struct {
	head *node
	tail *node
}

func (dll *Doubly) IsEmpty() bool {
	return dll.head == nil
}

func (dll *Doubly) Insert(data int) {
	newNode := node{data: data}

	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = &newNode
		dll.tail = &newNode
	}

}

func (dll *Doubly) Delete(data int) {
	if dll.IsEmpty() {
		fmt.Println("List is empty ")
	}
	if dll.head.data == data {

		dll.head = dll.head.next

	}
	curr := dll.head

	for curr != nil {
		if curr.data == data {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				dll.head = curr.next
			}

			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				dll.tail = curr.prev
			}

			fmt.Println("Delete data is ", curr.data)
			break

		}
		curr = curr.next
	}
	fmt.Println("No element data  matched ")
}

func (dll *Doubly) PrintAll() {

	if dll.IsEmpty() {
		fmt.Println("No data found ")
	}

	curr := dll.head

	for curr != nil {
		fmt.Println("Data is ", curr.data)
		curr = curr.next
	}

}
