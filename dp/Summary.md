### 背包问题

#### 背包递推公式

```javascript
//问能否能装满背包（或者最多装多少）
dp[j] = max(dp[j - nums[i]] + nums[i], dp[j]);

//问装满背包有几种方法
dp[j] = dp[j] + dp[j - nums[i]]

//问背包装满最大价值
dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);

//问装满背包所有物品的最小个数
dp[j] = min(dp[j - coins[i]] + 1, dp[j])
```

#### 背包和物品遍历顺序

- 01背包
    1. 二维数组 从小到大遍历 物品和背包先后顺序无所谓
    2. 一维数组 先遍历物品在遍历背包 背包从大到小遍历

- 完全背包
    1. 如果求组合数 *((1,2)和(2,1)是一样)* 就是外层for循环遍历物品，内层for遍历背包。
    2. 如果求排列数 *((1,2)和(2,1)是不一样)* 就是外层for遍历背包，内层for循环遍历物品。
    3. 如果求最小数 先后顺序无所谓

### 打家劫舍系列

1. 普通打家劫舍

```javascript
dp[i] = dp[i - 2] + nums[i], nums[i - 1]
```

2. 环形 分成两段取去处理

```javascript
[0, length - 1]
    [1, length]
```

3. 树形 后序遍历 核心定义节点的两个状态

```javascript
//抢当前node 则下家就不能抢了
node.Val + left[0] + right[0]
//不抢当前node 下家可以抢也可以不抢 取最大值
util.Max(left[0], left[1]) + util.Max(right[0], right[1])
```

### 股票买卖系列

### 子序列系列