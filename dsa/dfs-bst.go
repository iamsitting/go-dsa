package main

import "fmt"

// depth first search binary search tree
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

func search(curr *Node, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.value == needle {
		return true
	}

	// it only does one or the other
	if curr.value < needle {
		return search(curr.right, needle)
	} else {
		return search(curr.left, needle)
	}
}

func dfs(head *Node, needle int) bool {
	return search(head, needle)
}

func main() {
	n := &Node{value: 6}
	n.insert(3)
	n.insert(5)
	n.insert(10)
	n.insert(11)

	res := dfs(n, 10)
	fmt.Println("res: ", res)
}
