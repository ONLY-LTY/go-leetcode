package stack

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := &IHeap{}
	heap.Init(h)
	for _, value := range nums {
		heap.Push(h, value)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

type IHeap []int

func (h IHeap) Len() int { return len(h) }

func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *IHeap) Pop() any {
	//这里直接删除最后一个元素 heap包里面已经做了第一个元素和最后一个元素的替换
	//所以这里最后一个元素就是堆顶元素
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
