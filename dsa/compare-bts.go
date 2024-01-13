// compare the value and shape of two binary trees
package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if n == nil {
		return
	} else if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func compare(a *Node, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}

func main() {
	a := Node{value: 6}
	a.insert(3)
	a.insert(2)
	a.insert(7)
	a.insert(10)

	b := Node{value: 6}
	b.insert(3)
	b.insert(2)
	b.insert(7)
	b.insert(10)

	comp := compare(&a, &b)
	fmt.Println("is same:", comp)

	b.insert(2)

	comp = compare(&a, &b)
	fmt.Println("is same:", comp)
}
