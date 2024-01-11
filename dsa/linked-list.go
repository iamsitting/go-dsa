package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{value: data, next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ll *LinkedList) GetAll() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}

	fmt.Println()
}

func (ll *LinkedList) Pop() (int, error) {
	// there are no elements
	if ll.Head == nil {
		return 0, fmt.Errorf("no elements, can't be popped")
	}

	current := ll.Head

	for {
		if current.next.next == nil { // current is second to last element and current.next is the last node
			result := current.next.value
			current.next = nil
			return result, nil
		}
		current = current.next
	}
}

func main() {
	myList := LinkedList{}

	myList.Append(1)
	myList.Append(2)
	myList.Append(3)
	fmt.Print("Linked List:")
	myList.GetAll()

	r, _ := myList.Pop()
	fmt.Print("Popped", r)
	fmt.Print("Linked List:")
	myList.GetAll()

	r, _ = myList.Pop()

	fmt.Print("Popped", r)

	fmt.Print("Linked List:")
	myList.GetAll()
}
