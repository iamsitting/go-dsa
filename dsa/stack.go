package main

import "fmt"

type Node struct {
	value    int
	previous *Node // previous points from last to first
}

type Stack struct {
	length int
	head   *Node // head points to the "last" item in the stack
}

func (s *Stack) Peek() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.head.value, nil
}

func (s *Stack) Push(value int) {
	newNode := &Node{value: value, previous: nil}

	s.length++

	if s.head == nil {
		s.head = newNode
	} else {
		newNode.previous = s.head
		s.head = newNode
	}
}

func (s *Stack) Pop() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("stack is empty")
	}

	s.length--

	result := s.head.value

	if s.length == 0 { // popping last element
		s.head = nil
	} else {
		s.head = s.head.previous
	}

	return result, nil
}

func (s *Stack) DisplayAll() {
	current := s.head
	for i := 0; i < s.length; i++ {
		fmt.Print(current.value)
		current = current.previous
	}
	fmt.Println()
}

func main() {
	stack := &Stack{}

	fmt.Println("puhs 1,2,3,4")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.DisplayAll()

	fmt.Println("pop")
	stack.Pop()
	stack.Peek()

	stack.DisplayAll()

	fmt.Println("pop")
	stack.Pop()

	fmt.Println("pop")
	stack.Pop()

	stack.DisplayAll()
}
