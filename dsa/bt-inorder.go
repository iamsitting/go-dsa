package main

// tree traversal - preorder

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

func walk(curr *Node, path []int) []int {
	if curr == nil {
		return path
	}

	path = walk(curr.left, path)
	path = append(path, curr.value)
	path = walk(curr.right, path)

	return path
}

func in_order_search(head *Node) []int {
	entry := []int{}
	result := walk(head, entry)
	fmt.Println(result)
	return result
}

func main() {
	node := &Node{value: 10}
	node.insert(2)
	node.insert(11)
	node.insert(3)
	node.insert(4)
	node.insert(5)
	node.insert(6)
	in_order_search(node)
}
