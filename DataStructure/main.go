package main

import (
	"DataStructure/tree"
)

func main() {
	// linked := linkedlist.LinkedList{}

	// linked.Push(10)
	// linked.Push(20)
	// linked.Push(30)
	// linked.Push(40)
	// linked.Push(50)
	// linked.Push(60)
	// linked.Push(70)

	// linked.GetMiddle()

	// fmt.Println(" Display Data is ******  ")
	// linked.Display()

	// fmt.Println(" Pop Operation ")
	// linked.Pop()

	// fmt.Println(" Display Data is ******  ")
	// linked.Display()

	// fmt.Println("  @@@@@@  After Reverse @@@@@")
	// linked.Reverse()
	// linked.Display()

	// fmt.Println(" Rotate ++++++++  ")
	// linked.RotateWithHead(3)
	// linked.Display()

	// fmt.Println(" Working on Stack ")

	// stack := stack.Stack{}

	// stack.Push(10)
	// stack.Push(20)
	// stack.Push(30)

	// stack.PrintAll()

	// stack.Pop()
	// stack.PrintAll()

	// fmt.Println(" Top value is ", stack.Peek())

	// fmt.Println(" +++++++++++  Queue ++============ ")

	// queue := quque.Queue{}

	// queue.Enqueue(10)
	// queue.Enqueue(20)
	// queue.Enqueue(30)

	// queue.PrintAll()

	// queue.Dequeue()

	// queue.PrintAll()

	// fmt.Println("  ========  Queue ============  ")

	// dl := doublelinkedlist.Doubly{}
	// dl.PrintAll()

	// dl.Insert(10)
	// dl.Insert(20)
	// dl.Insert(30)

	// dl.PrintAll()

	// dl.Delete(50)
	// dl.Delete(30)

	// dl.PrintAll()

	// fmt.Println("++++++++++++  Working on Heap +++++++++++++++   ")
	// hp := h.MinHeap{9, 6, 2, 4, 3, 1, 5, 6}

	// heap.Init(&hp)

	// fmt.Println("Heap is ", hp)
	// fmt.Println(" Before  adder ", &hp)
	// heap.Push(&hp, 10)
	// fmt.Println(" After adder ", &hp)

	// fmt.Println(" Pop is ", &hp, hp.Pop())

	// fmt.Println("^^^^^^^^^^^^^  Tree ^^^^^^^^^^^^^^ ")

	tree := tree.Node{Data: 10}
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)

	tree.Insert(50)
	tree.TraverseNode()

}
