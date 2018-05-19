package binaryHeap

import (
	"testing"
)

func assertSlicesEqual(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

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
	exp := []int{2, 3}

	for i := 3; i > 0; i-- {
		bh.Insert(i)
	}
	t.Log("Before polling:", bh.Nodes)
	bh.Poll()
	t.Log("After polling:", bh.Nodes)

	eq := assertSlicesEqual(bh.Nodes, exp)

	if !eq {
		t.Error("Incorrect Polling")
	}
}
