package binaryHeap

import (
	"fmt"
)

type BinaryHeapMethods interface {
	Insert(item int)
	GetParentIndex(i int) int
}

type BinaryHeap struct {
	Nodes []int
	Size  int
}

func (b *BinaryHeap) Peek() int {
	return b.Nodes[0]
}

func (b *BinaryHeap) Insert(item int) {
	b.Nodes = append(b.Nodes, item)
	i := len(b.Nodes) - 1
	p := b.GetParentIndex(i)
	b.Size++

	for b.Nodes[i] < b.Nodes[p] {
		b.Nodes[i], b.Nodes[p] = b.Nodes[p], b.Nodes[i]
		i = p
		p = b.GetParentIndex(i)
	}
}

func (b *BinaryHeap) Poll() int {
	root := b.Nodes[0]

	b.Nodes[0] = b.Nodes[b.Size-1]
	b.Size--

	i := 0

	for b.HasLeftChild(i) {
		smallerChild := &b.Nodes[i*2+1]
		smindx := i*2 + 1

		if *smallerChild > b.Nodes[i*2+2] {
			smallerChild = &b.Nodes[i*2+2]
			smindx = i*2 + 2
		}

		if *smallerChild > b.Nodes[i] {
			break
		}
		*smallerChild, b.Nodes[i] = b.Nodes[i], *smallerChild
		i = smindx
	}

	return root
}

func (b *BinaryHeap) GetParentIndex(i int) int {
	return (i - 1) / 2
}

func (b *BinaryHeap) HasLeftChild(i int) bool {
	cid := i*2 + 1
	return cid < b.Size
}

func (b *BinaryHeap) HasRightChild(i int) bool {
	cid := i*2 + 2
	return cid < b.Size
}

func main() {
	fmt.Println("haha")
}
