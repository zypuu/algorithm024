学习笔记

## 深度优先搜索-dfs

通过递归进行深度优先

``` javascript
// 二叉树
func dfs(node *TreeNode) {
	// 当前节点已经访问过，则return
	if _, ok := visited[node]; ok {
 		return
	}
 	// 未访问则记录,处理当前层
 	visited[node] = true
	// 左右子树，继续向下深度遍历
	dfs(node.Left)
	dfs(node.Right)
	
}

// 多叉树，另一种形式
func dfs(node *TreeNode) {
 	// 未访问则记录,处理当前层
 	visited[node] = true
	// 左右子树，继续向下深度遍历
	for _, v := range node.Children {
		//没在访问记录里，则进入下一层递归
		if _, ok := visited[node]; !ok {
	 		dfs(v)
		}
	}
	
}

// 非递归写法,手动维护一个栈，栈与递归后进先出
func dfs(root *TreeNode) {
	// root不存在则返回
	if root == nil {
		return
	}
	// 手动维护一个栈，初始化，并放入root，还有一个visited的map
	stack := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
 	// 未访问则记录,处理当前层
 	for len(stack) > 0 {
 		// 节点出栈
 		node := stack[len(stack) - 1]
 		stack = stack[:len(stack) - 1]
 		// 访问记录
 		visited[node] = true
 		// 处理逻辑
 		process(node)
 		// 更新node，或更新栈
 		nodes = related(node)
 		stack = append(stack, nodes...)
 		
 	}
	
}

```

## 广度优先搜搜-BFS

通过队列实现，先进先出，一层一层遍历

``` javascript
// 广度优先，队列遍历
func bfs(root *TreeNode) {
	// root不存在则返回
	if root == nil {
		return
	}
	// 手动维护一个队列，初始化，并放入root，还有一个visited的map
	q := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
 	// 未访问则记录,处理当前层
 	for len(q) > 0 {
 		// 节点出队
 		node := q[0]
 		q = q[1:]
 		// 访问记录
 		visited[node] = true
 		// 处理逻辑
 		process(node)
 		// 更新q，更新node
 		nodes = related(node)
 		q = append(q, nodes...)
 	}
	
}

```

两种搜索的方式区别：

dfs类似入栈，先进后出

bfs类似入队，先进先出

## 贪心算法

每一步选取当前最优解

问题能够分解成子问题，子问题最优解，递推到最终最优解

与动态规划区别：

贪心算法：每个子情况都进行选择，不能回退

动态规划：保存以前的运算结果，可以回退

PS： 关键点在于如何证明可以使用贪心算法，能得到想要的结果

## 二分查找

二分查找条件：

1.目标函数单调性，递增或者递减

2.存在上下界

3.能够通过索引访问

``` javascript
// 二分查找
func two(nums []int) int {
	// 左右起点
	left, right := 0, len(nums) - 1
	for left <= right {
		// 得到中间值
		mid = (left + right) / 2
		// 等于找到则返回
		if nums[mid] == target {
			return target or break
		} else if mid < target { // 否则target在mid右侧，左下标到mid+1
			left = mid + 1
		} else { // target在mid左侧，右下标到mid-1
			right = mid - 1
		}
	}


```
每次循环将搜索范围缩小到之前的一半， O（logN）

### 半有序数组

使用二分查找，寻找一个半有序数组中间无序的地方[4,5,6,7,0,1,2]

``` javascript
// 二分查找，返回的是数组的索引
func search(nums []int) int {
	// 边界条件
    if len(nums) < 2 {
        return 0
    }
    // 左右下标
    l := 0
    r := len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        // 最后左右相等的时候就是转折
        if l == r {
            return l
        }
        // 如果左边小于中间值，左半部分递增，则转折点在右边，否在在左边
        if nums[r] < nums[mid] {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return 0
}

// 暴力法
func search(nums []int) int {
	// 边界条件
    if len(nums) < 2 {
        return 0
    }
    for i, v := range nums {
    	// 出现转折，说明下一个值小于前一个值了
        if v > nums[i+1] {
            return i
        }
    }
    return 0
}
```

该题类似搜索旋转数组，搜索旋转数组，是要寻找目标值，则在二分查找时，进行数值比对，根据比对不同情况，对左右下标的不同处理

### 有效的完全平方数

二分查找方法

``` javascript
// 二分查找
func isPerfectSquare(num int) bool {
	// 边界条件
    if num < 2 {
        return true
    }
    l := 0
    // 平方根一定在一半之前，后面的平方肯定大于num
    r := num / 2
    for l <= r {
        mid := (l + r) / 2
        // 找到结果返回
        if mid * mid == num { return true }
        if mid * mid > num { 
            r = mid - 1
        } else {l = mid + 1}
    }
    return false
}
```

奇数方法
完全平方数都是由奇数组成的，每次减去奇数，最后是0则是完全平方数

``` javascript
func isPerfectSquare(num int) bool {
    n := 1
    for num > 0 {
        num -= n
        n += 2
    }
    return num == 0
}
```

牛顿迭代

以当前点的切线斜率公式推导 x^2 - num = 0

当前点的切线交于x轴为y点

f(x)/(x - y) = f`(x)

(x^2 - num)/(x-y) = 2x

x^2 - num = 2x^2-2xy

y=(x+num/x)/2

``` javascript
func isPerfectSquare(num int) bool {
    if num < 2 {
        return true
    }
    // 从一半开始
    x := num / 2
    // 不断迭代，找其切线斜率的坐标，直到，满足表达式x^2 - num = 0
    for x * x > num {
        x = (x + num / x) / 2
    }
    return x * x == num
}
```


