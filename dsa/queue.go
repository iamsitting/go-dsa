package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) Peek() (int, error) {
	if q.head == nil {
		return 0, fmt.Errorf("This queue has no elements")
	}

	return q.head.value, nil
}

func (q *Queue) Deque() (int, error) {
	if q.head == nil {
		return 0, fmt.Errorf("This queue has no elements")
	}

	q.length--

	temp := q.head.value
	q.head = q.head.next
	return temp, nil
}

func (q *Queue) Enqueue(value int) {
	q.length++

	newNode := &Node{value: value, next: nil}

	if q.tail == nil {
		q.tail = newNode
		q.head = newNode
		return
	}

	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue) DisplayAll() {
	fmt.Print("All:")
	current := q.head
	for i := 0; i < q.length; i++ {
		fmt.Print(current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	q := Queue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	q.DisplayAll()

	r, _ := q.Peek()

	r, _ = q.Deque()

	fmt.Println("deq: ", r)

	q.DisplayAll()

	r, _ = q.Peek()
	fmt.Println("peek: ", r)

	q.Deque()
	q.DisplayAll()
}
