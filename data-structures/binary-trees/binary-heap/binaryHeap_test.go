package binaryHeap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	bh := BinaryHeap{[]int{}, 0}
	bh.Insert(0)

	if len(bh.Nodes) == 0 {
		t.Error("Binary Heap isn't inserting items")
	}
}

func TestGetParentIndex(t *testing.T) {
	bh := BinaryHeap{[]int{}, 0}

	zero := bh.GetParentIndex(0)
	one := bh.GetParentIndex(1)
	eight := bh.GetParentIndex(8)

	if zero != 0 {
		t.Error("Doesn't handle top element")
	}

	if one != 0 {
		t.Error("Doesn't hanlde odd children")
	}

	if eight != 3 {
		t.Error("Doesn't hanlde even children")
	}
}
func TestCorrectInsertion(t *testing.T) {
	bh := BinaryHeap{[]int{}, 0}

	for i := 10; i > 0; i-- {
		bh.Insert(i)
	}

	for i := 0; i < len(bh.Nodes); i++ {
		p := bh.GetParentIndex(i)
		parent := bh.Nodes[p]
		child := bh.Nodes[i]

		if parent > child {
			t.Error("Parent is greater than child!")
		}
	}
}

func TestPoll(t *testing.T) {
	bh := BinaryHeap{[]int{}, 0}

	for i := 10; i > 0; i-- {
		bh.Insert(i)
	}

	bh.Poll()
	if true {
		t.Error("After polling:", bh.Nodes)
	}
}
