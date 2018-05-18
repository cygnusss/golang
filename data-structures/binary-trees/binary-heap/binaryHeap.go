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
}

func (b *BinaryHeap) Peek() int {
	return b.Nodes[0]
}

func (b *BinaryHeap) Insert(item int) {
	b.Nodes = append(b.Nodes, item)
	i := len(b.Nodes) - 1
	p := b.GetParentIndex(i)

	for b.Nodes[i] < b.Nodes[p] {
		b.Nodes[i], b.Nodes[p] = b.Nodes[p], b.Nodes[i]
		i = p
		p = b.GetParentIndex(i)
	}
}

func (b *BinaryHeap) GetParentIndex(i int) int {
	return (i - 1) / 2
}

func main() {
	fmt.Println("haha")
}
