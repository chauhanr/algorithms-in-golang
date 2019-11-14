package ch1

import (
	"container/heap"
	"testing"
)

/*TestIntegerHeap Method to test the heap function*/
func TestIntegerHeap(t *testing.T) {
	result := []int{1, 2, 4, 5}

	iheap := &IntegerHeap{1, 4, 5}
	heap.Init(iheap)
	heap.Push(iheap, 2)

	i := 0
	t.Logf("min: %d\n", (*iheap)[0])
	for iheap.Len() > 0 {
		v := heap.Pop(iheap).(int)
		if v != result[i] {
			t.Errorf("Error result expected %d but for %d\n", result[i], v)
		}
		i++
	}

}
