package ch1

/*IntegerHeap Declarting a heap type Integer*/
type IntegerHeap []int

func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

/*Push will add a new element into the heap and rearrange the heap*/
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

/*Pop the first element on the heap reheap the heap*/
func (iheap *IntegerHeap) Pop() interface{} {
	var n, x1 int
	var previous = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}
