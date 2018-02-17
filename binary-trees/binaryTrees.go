package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

func (h *node) Insert(v int) *node {

	newNode := new(node)
	newNode.val = v
	temp := h

	for {

		if v < temp.val {
			if temp.left == nil {
				temp.left = newNode
				return temp.left
			}
			temp = temp.left
		} else if v >= temp.val {
			if temp.right == nil {
				temp.right = newNode
				return temp.right
			}
			temp = temp.right
		}

	}

}

func main() {

	var head *node
	head = new(node)

	head.val = 4
	fmt.Printf("Node is: %v\n", head)

	head.Insert(5)
	fmt.Printf("Right node is: %v\n", head.right)

	head.Insert(3)
	fmt.Printf("Left node is: %v\n", head.left)

}
