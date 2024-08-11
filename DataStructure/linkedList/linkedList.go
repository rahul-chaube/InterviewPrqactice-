package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (node *LinkedList) Push(data int) {

	newNode := &Node{Data: data, Next: nil}

	if node.head == nil {
		node.head = newNode
	} else {
		currentNode := node.head

		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}

}

func (node *LinkedList) Pop() {

	if node == nil {
		fmt.Println("EmptyNode")
	} else {
		currentNode := node.head
		preNode := node.head

		for currentNode.Next != nil {
			preNode = currentNode
			currentNode = currentNode.Next
		}
		fmt.Println("Pop Data is ", currentNode.Data)
		preNode.Next = nil
	}

}

func (node *LinkedList) Display() {

	if node == nil {
		fmt.Println("Empty Node ")
	} else {
		currentNode := node.head

		for {
			fmt.Println("Node data is ", currentNode.Data)

			if currentNode.Next == nil {
				break
			}
			currentNode = currentNode.Next

		}
	}

}

func (node *LinkedList) GetMiddle() {
	var counter int = 1
	if node.head == nil {
		fmt.Println("Invalid Node ")
	}
	var midleNode *Node

	currentNode := node.head
	midleNode = node.head

	for currentNode != nil {

		if counter%2 == 0 {
			midleNode = midleNode.Next
		}
		currentNode = currentNode.Next
		counter++

	}

	fmt.Println("Middle node data is ", midleNode.Data)

}

func (node *LinkedList) Reverse() {

	// Initializa node to prev,next and current node

	var prev, next *Node

	curr := node.head

	for curr != nil {
		next = curr.Next // Next node will contain next with current, In first node will be
		curr.Next = prev // assign current node with next prev will became first node {data:<data>,node:nil} II node
		prev = curr      //make current node is previous, firnode will be nill prev:={data:<data>,node:nil}

		curr = next // make current node is next node for itterate next node
	}
	node.head = prev
}

func (node *LinkedList) RotateWithHead(k int) {

	if node.head == nil {
		return
	}
	count := 1

	current := node.head

	for count < k && current != nil {
		current = current.Next
		count++
	}

	if current == nil {
		fmt.Println("Current is last node")
		return
	}

	knode := current

	for current.Next != nil {
		current = current.Next
	}

	fmt.Println("Current node is ", current.Data, current.Next)

	current.Next = node.head

	node.head = knode.Next
	knode.Next = nil

}
