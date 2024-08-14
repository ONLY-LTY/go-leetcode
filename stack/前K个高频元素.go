package stack

import "container/heap"

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 示例 1:
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:
//
// 输入: nums = [1], k = 1
// 输出: [1]
//
// 提示：
//
// 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小
func topKFrequent(nums []int, k int) []int {
	frequent := make(map[int]int)
	for _, n := range nums {
		frequent[n]++
	}
	h := &IntHeap{}
	heap.Init(h)
	for key, value := range frequent {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// IntHeap 小顶堆
type IntHeap [][2]int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.([2]int)) }

func (h *IntHeap) Pop() any {
	//这里直接删除最后一个元素 heap包里面已经做了第一个元素和最后一个元素的替换
	//所以这里最后一个元素就是堆顶元素
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
