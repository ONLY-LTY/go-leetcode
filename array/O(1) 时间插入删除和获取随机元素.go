package array

import "math/rand"

// 实现RandomizedSet 类：
//
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
//
// 示例：
//
// 输入
// ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
// [[], [1], [2], [2], [], [1], [2], []]
// 输出
// [null, true, false, true, 2, true, false, 2]
//
// 解释
// RandomizedSet randomizedSet = new RandomizedSet();
// randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
// randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
// randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
// randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
// randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
// randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
// randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。

type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{indices: make(map[int]int)}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.indices[val]; ok {
		return false
	}
	r.nums = append(r.nums, val)
	r.indices[val] = len(r.nums) - 1
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.indices[val]; !ok {
		return false
	}
	//删除元素index
	rIndex := r.indices[val]
	lastVal := r.nums[len(r.nums)-1]
	//最后一个元素移动到删除元素的位置
	r.nums[rIndex] = lastVal
	//删除最后一个元素
	r.nums = r.nums[:len(r.nums)-1]
	//map中更新最后一个元素的索引
	r.indices[lastVal] = rIndex
	//map中删除元素val
	delete(r.indices, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
