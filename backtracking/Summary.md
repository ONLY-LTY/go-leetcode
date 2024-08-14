>*回溯法抽象为树形结构后，其遍历过程就是：for循环横向遍历，递归纵向遍历，回溯不断调整结果集。*


#### 1. 元素无重不可复选
即 nums 中的元素都是唯一的，每个元素最多只能被使用一次，backtrack 核心代码如下：
```go
// 组合/子集问题回溯算法框架 
func backtrack(nums []int, start int) {
	// 回溯算法标准框架 
	for i := start; i < len(nums); i++ {
		// 做选择
		track = append(track, nums[i])
		// 注意参数 
		backtrack(nums, i+1)
		// 撤销选择 
		track = track[:len(track)-1]
	}
}

// 排列问题回溯算法框架
func backtrack(nums []int) {
	for i := 0; i < len(nums); i++ {
		// 剪枝逻辑 
		if used[i] {
			continue
		}
		// 做选择 
		used[i] = true
		track = append(track, nums[i])

		backtrack(nums)
		// 撤销选择 
		track = track[:len(track)-1]
		used[i] = false
	}
}

```
#### 2. 元素可重不可复选
即 nums 中的元素可以存在重复，每个元素最多只能被使用一次，其关键在于排序和剪枝，backtrack 核心代码如下：
```go
import "sort"

/* 组合/子集问题回溯算法框架 */
func backtrack(nums []int, start int) {
    // 回溯算法标准框架
    for i := start; i < len(nums); i++ {
        // 剪枝逻辑，跳过值相同的相邻树枝
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        // 做选择
        track = append(track, nums[i])
        // 注意参数
        backtrack(nums, i+1)
        // 撤销选择
        track = track[:len(track)-1]
    }
}

func combinations(nums []int) {
    // 预处理排序
    sort.Ints(nums)
    backtrack(nums, 0)
}

/* 排列问题回溯算法框架 */
func backtrack(nums []int) {
    for i := 0; i < len(nums); i++ {
        // 剪枝逻辑
        if used[i] {
            continue
        }
        // 剪枝逻辑，固定相同的元素在排列中的相对位置
        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
            continue
        }
        // 做选择
        used[i] = true
        track = append(track, nums[i])

        backtrack(nums)

        // 撤销选择
        track = track[:len(track)-1]
        used[i] = false
    }
}
```
#### 3. 元素无重可复选
即 nums 中的元素都是唯一的，每个元素可以被使用若干次，只要删掉去重逻辑即可，backtrack 核心代码如下：
```go
/* 组合/子集问题回溯算法框架 */
func backtrack(nums []int, start int) {
    // 回溯算法标准框架
    for i := start; i < len(nums); i++ {
        // 做选择
        track = append(track, nums[i])
        // 注意参数
        backtrack(nums, i)
        // 撤销选择
        track = track[:len(track)-1]
    }
}
/* 排列问题回溯算法框架 */
func backtrack(nums []int) {
    for i := 0; i < len(nums); i++ {
        // 做选择
        track = append(track, nums[i])
        backtrack(nums)
        // 撤销选择
        track = track[:len(track)-1]
    }
}

```