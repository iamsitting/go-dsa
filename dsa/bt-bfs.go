package main

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

func deque(q []*Node) (*Node, []*Node) {
	return q[0], q[1:]
}

func bfs(node *Node, needle int) bool {
	q := []*Node{node}

	for len(q) > 0 {
		next, q := deque(q)

		if next == nil {
			continue
		}
		if next.value == needle {
			return true
		}

		q = append(q, next.left)
		q = append(q, next.right)
	}
	return false
}
