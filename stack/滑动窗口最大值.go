package stack

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。
func maxSlidingWindow(nums []int, k int) []int {
	q := newQueue()
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	res = append(res, q.Front())

	for i := k; i < len(nums); i++ {
		//移除窗口最前面的值
		q.Pop(nums[i-k])
		//滑动窗口
		q.Push(nums[i])
		res = append(res, q.Front())
	}
	return res
}

// 单调队列
func newQueue() *MonotoneQueue {
	return &MonotoneQueue{queue: make([]int, 0)}
}

type MonotoneQueue struct {
	queue []int
}

func (q *MonotoneQueue) Front() int {
	return q.queue[0]
}

func (q *MonotoneQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *MonotoneQueue) Push(Val int) {
	for len(q.queue) > 0 && Val > q.Back() {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, Val)
}

func (q *MonotoneQueue) Pop(Val int) {
	if len(q.queue) > 0 && Val == q.Front() {
		q.queue = q.queue[1:]
	}
}
