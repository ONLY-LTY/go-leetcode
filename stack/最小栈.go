package stack

import "go-leetcode/util"

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。

type MinStack struct {
	//栈元素
	element []int
	//栈最小元素 辅助栈
	mm []MaxMin
}
type MaxMin struct {
	max int
	min int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]MaxMin, 0)}
}

func (s *MinStack) Push(val int) {
	//入栈
	s.element = append(s.element, val)
	if len(s.mm) == 0 {
		s.mm = append(s.mm, MaxMin{max: val, min: val})
	} else {
		//最小元素入最小栈
		top := s.mm[len(s.mm)-1]
		s.mm = append(s.mm, MaxMin{min: util.Min(val, top.min), max: util.Max(val, top.max)})
	}
}

func (s *MinStack) Pop() {
	//出栈
	s.element = s.element[:len(s.element)-1]
	//最小元素出栈
	s.mm = s.mm[:len(s.mm)-1]
}

func (s *MinStack) Top() int {
	return s.element[len(s.element)-1]
}

func (s *MinStack) GetMin() int {
	return s.mm[len(s.mm)-1].min
}
func (s *MinStack) GetMax() int {
	return s.mm[len(s.mm)-1].max
}
