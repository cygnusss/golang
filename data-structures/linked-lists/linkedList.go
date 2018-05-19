package main

import (
	"fmt"
)

// LinkedListMethods ...
type LinkedListMethods interface {
	Insert(interface{})
	Reverse()
	Print()
}

// LinkedList creates a linked list
type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
}

// LinkedListNode creates a node in the linked list
type LinkedListNode struct {
	Val      interface{}
	NextNode *LinkedListNode
}

// Insert inserts a linked list
func (c *LinkedList) Insert(value interface{}) {
	node := LinkedListNode{value, nil}

	if c.Head == nil {
		c.Head = &node
	}
	if c.Tail != nil {
		c.Tail.NextNode = &node
	}
	c.Tail = &node
}

// Reverse reverses a linked list
func (c *LinkedList) Reverse() {
	var prev *LinkedListNode
	var next *LinkedListNode
	curr := c.Head

	for curr != nil {
		next = curr.NextNode
		curr.NextNode = prev
		prev = curr
		curr = next
	}

	c.Tail, c.Head = c.Head, c.Tail

	_, _ = prev, next
}

// Print prints all nodes in a given linked list
func (c *LinkedList) Print() {
	root := c.Head
	for root != nil {
		fmt.Print(root.Val, "->")
		root = root.NextNode
	}
	fmt.Printf("nil\n")
}
