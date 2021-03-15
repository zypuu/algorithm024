学习笔记

## 动态规划思路

1、根据大问题找重复子问题

2、定义dp状态数组，明确dp[i]含义

3、找dp方程，明确边界条件等，不同处理


## 动态规划解题步骤

### 判断边界条件

1、给的值为空，提前返回

2、长度为1这种特殊值

### 初始化dp

#### 初始化一维dp数组， dp=[]，要明白dp[i]的意义
初始化方式: 
	新开数组，并赋上初始0值，然后迭代修改
	以原数组直接为dp（优化空间复杂度，不过会修改原数组的值）
例：
	1、爬楼梯问题：
		新开数组，dp[i]代表，第i个数的斐波那契数值，最后迭代到dp[i]即可
		初始化dp[0]=1,dp[1]=1,第一个数是1，0是用于初始计算，n是正整数
	2、打家劫舍问题：
		新开dp数组，dp[i]表示当前的最大值，初始化0值，1值（要不偷0，要不偷1，两个的最大值）
		不开dp数组，直接nums为dp，初始化nums[1]=max(nums[0], nums[1])
	3、零钱兑换问题：
		新开dp，初始化一个amount+1的数组，dp[i]组成该金额的最小组合，最终结果是dp[amount]
		金额依次递增，算出每个金额的最小组合，索引即代表金额



#### 一维dp数组通过优化，可将空间复杂度降为1，通过指针迭代
要明白每个指针的意义，以及特殊问题相关
一般都会2个-3个指针，分别代表，现在的值，前一个值，结果值或后一个值，迭代过程中不断更新指针值
例：
	1、爬楼梯问题：
		对于上面开数组的优化， a,b,c 三个指针，n是正整数，初始化a=0，b=1，c=1
		a：n-2， b：n-1， c：n
	2、打家劫舍问题：
		数组优化，两个指针cur,pre,pre前一个，cur当前个
	3、最大子序和问题：
		不需要初始化数组，根据特殊情况，双指针迭代，或者单指针，修改数组，循环中舍弃不符合条件的值
		res最终结果，sum当前和，res初始化数组第一个值，从第二个开始
	4、最大乘积和：
		初始化一个最大值，一个最小值，结果，三个值，处理正负数的问题
	5、解码方法问题：
		初始化两个指针，一个当前位置的解码方法，一个前一个位置的解码方法


#### 初始化二维dp数组，dp=[][]int, dp[i][j]的意义
例：
	1、打家劫舍问题（一维问题二维化）：
		初始化dp数组，dp=[][]int,升维思路记录额外数据，dp[i][0,1],0代表不偷，1代表偷
		并初始化0值,dp[0][0],dp[0][1],初始值偷与不偷的情况
	2、三角形最小路径和问题：
		初始化dp数组[][]int,并将最后一行赋值，逆向循环，从倒数第二行开始
		不开dp数组，直接triangle，不需要赋值了
		优化空间复杂度，只需dp[],初始化最后一层，记录每一层，逆向循环，整层更新掉
	3、回文子串（一维问题二维化）：
		初始化二维dp数组，记录i，j下标为起始的字符串是否是回文串，true，false
		可优化为1维，在一次里层循环中使用，外层直接从头更新
	4、最大正方形问题：
		初始化二维数组，并赋值1，与原数组对应，dp[i][j]代表最大正方形的边长，并初始化最初结果1
	5、最小路径和问题：
		初始化二维dp，dp[i][j]代表到当前点的最小路径和
	6、编辑距离问题：
		初始化二维dp，长度加1，单词前面加一个空串处理，dp[i][j]表示，单词1和单词2的前ij字符匹配所需要的最小编辑次数


### 循环迭代dp值

根据循环方向不同，循环初始值不同，结束值不同，维度不同分多这种情况

每次循环处理dp方程,区分情况，注意边界条件，分治，可能各个条件的dp方程不一样

注意dp[i]的意义，根据循环的正序，倒序返回

#### 一维循环：从开始往后循环，i++

例：
	1、爬楼梯问题:
		初始值，初始到了dp[1]，从2开始循环到n，返回dp[n]，初始的数组要到n+1
		dp方程：dp[i] = dp[i-1] + dp[i-2]
		返回dp[n]
		指针形式：c就相当于n的值，初始化了1的值，从1开始循环到n
		先更新值，然后计算
		a = b
		b = c
		c = a + b
		返回c
	2、打家劫舍问题：（打家劫舍2就是分两种情况dp）
		初始化是二维数组：
		dp[i]每次计算偷与不偷的情况，最后返回偷与不偷的最大值，从i=1开始，正向
		dp方程：	dp[i][0] = max(dp[i-1][0], dp[i-1][1]) // i不偷，i-1偷与不偷之间的最大值
        		dp[i][1] = nums[i] + dp[i-1][0] // i偷，则当前i加上i-1不偷的值
        初始化是一维数组：
        dp[i]是代表最大值，不关心偷与不偷，只获取偷和，不偷的最大值， i-2包含了i-3，i-4等
        dp方程：dp[i] = max(dp[i-1], dp[i-2] + nums[i])
        不开dp数组，直接nums为dp 初始化nums[1]： 
        dp方程：nums[i] = max(nums[i-1], nums[i-2] + nums[i])
        指针pre          ，cur从2开始循环：
        dp方程：pre, cur = cur, max(cur, pre + nums[i])
    3、最大子序和：
    	res记录结果，sum当前值，然后对sum进行累加，如果sum累加之后小于0，则把从下一个值从新开始，赋值为sum
    	比较结果，更新res
    	另一种思路：更新nums[i],进行累加nums[i] + nums[i-1] > nums[i]，说明nums[i-1]是正数，否则nums[i-1]是负数
    	则不动，nums[i]就是dp[i]代表累加和
    4、最大乘积和：
    	每次循环，获取当前的最大，最小值，每次更新最大最小值 
    	最大值：比较最大值*当前值，当前值，当前值*最小值的大小，针对0，变号的处理
    	最小值：比较最小值*当前值，当前值，当前值*最大值的大小，针对0，变号的处理
    	另外思路：负数的个数，偶数则是全部最大，奇数则从前，从后连乘，舍弃1个负数的两种情况的最大值
    5、解码方法问题：
    	先根据0的位置分情况，只有10,20符合且只能这么解码，所以次数不变
    	除了上一种其他0都不符合
    	然后找正常可组成26之内的，则解码方法等于加上之前的，然后更新cur，pre
    	否则就是单数字那种，更新pre


#### 一维循环：从后往前循环，i--

#### 二维循环：从前往后循环，i--

例：
	1、零钱兑换问题：
    	外层循环：从第一个金额开始向后遍历，每个金额初始dp[i] = -1
    	内层循环：遍历硬币数组当前金额即索引小于硬币数，则为-1跳过或者减去这个硬币金额后前一个也无解是-1
    		 前一个有解的话，就是前一个所需的最小组合数+1，如果这个数字小于dp[i],或者dp[i]没有计算过,则更新dp[i]
   	最后返回amount索引的金额
   	2、回文串问题：
   		外层循环：字符串结束下标
   		里层循环：字符串起始下标，记录起始下标到结束下标（3种情况：1，ij相等，2：ij差1且相等，3，ij差2以上且相等且里面的dp是true，是回文串），一维优化后，只记录起始下标i，不符合记录false，下一个外层从头更新
   	3、最大正方形问题：
   		以右下角包含1的为正方形的右下角，其他三个点的最小值+1
   		dp方程：dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
   	4、最小路径和问题：
   		分边界考虑：边界：没得选择，加上前面那点的dp值，加上原数组的值
   		非边界：前面的两个的最小值，加上原数组值
   		dp方程：dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
   	5、编辑距离问题：
   		里外层循环，两个单词两个指针位置移动，从0开始，0是空串，边界条件初始化（一个单词为空串）
   		当word ij相等时，不需要操作，ij = i-1，j-1
   		word ij不等时，三种操作下的最小值+1
   		替换操作：i-1与j-1匹配，ij替换即可  +1
   		插入操作：i与j-1匹配，i后面插入j  +1
   		删除操作：i-1与j匹配，删除i  +1
   		dp方程： if word1[i-1] == word2[j-1]: dp[i][j] = dp[i - 1][j - 1]
                else：dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
        最后返回dp[m][n]


#### 二维循环：从后往前循环，i--

自下向上循环，注意初始值的处理，看看最后一层是不是需要进入循环，不进入直接初始化最后一层

例：
	1、最小三角形路径和：
	外层循环从倒数第二层开始。n-2往前，里层循环正向，dp[i][j]为下一层相邻的最小值
	dp方程：dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
	优化空间复杂度：dp[]初始化最后一层，向上循环，每次整层更新
	dp方程：dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]


## 动态规划例题

### 爬楼梯

``` javascript
func climbStairs(n int) int {
    a, b, c := 0, 1, 1
    for i := 2; i <= n; i++ {
        a = b
        b = c
        c = a + b
    }
    return c
}
```

### 打家劫舍

``` javascript
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }
    pre := nums[0]
    cur := max(nums[0], nums[1])
    for i := 2; i < n; i++ {
        pre, cur = cur, max(cur, pre + nums[i])
    }
    return cur
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

### 打家劫舍2


``` javascript
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = nums[0]
    for i := 2; i < n - 1; i++  {
        dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
    }
    res1 := dp[n - 2]
    dp[0] = 0
    dp[1] = nums[1]
    for i := 2; i < n; i++  {
        dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
    }
    return max(res1, dp[n-1])
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

### 零钱兑换

完全背包问题

``` javascript
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := 1; i <= amount; i++ {
        dp[i] = -1
        for _, c := range coins {
            if i < c || dp[i-c] == -1 {
                    continue
            }
            
            count := dp[i-c] + 1
            if dp[i] == -1 || dp[i] > count {
                    dp[i] = count
            }
        }
    }
    return dp[amount]
}
```

### 最大子序和


``` javascript
func maxSubArray(nums []int) int {
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] + nums[i-1] > nums[i]{
            nums[i] += nums[i -1]
        }
        if nums[i] > res {
            res = nums[i]
        }
        
    }
    return res
}
```

### 最大乘积

遍历法，根据负数的数量是奇数还是偶数个

``` javascript
func maxProduct(nums []int) int {
    res := nums[0]

    cur := 1
    for i := 0; i < len(nums); i++ {
        cur *= nums[i]
        res = max(cur, res)
        if nums[i] == 0 {cur = 1}
    }
    cur = 1
    for i := len(nums) - 1; i >= 0; i-- {
        cur *= nums[i]
        res = max(cur, res)
        if nums[i] == 0 {cur = 1}
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

动态规划


``` javascript
func maxProduct(nums []int) int {
    mxf, mnf, res := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        mx, mn := mxf, mnf
        mxf = max(mx * nums[i], max(nums[i], nums[i] * mn))
        mnf = min(mn * nums[i], min(nums[i], nums[i] * mx))
        res = max(mxf, res)
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
```

### 三角形最小路径和

``` javascript
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    dp := make([]int, len(triangle[n-1]))
    for i := 0; i < len(triangle[n-1]); i++ {
        dp[i] = triangle[n-1][i]
    }
    for i := n - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
        }
    }
    return dp[0]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```

### 回文串

``` javascript
func countSubstrings(s string) int {
    n := len(s)
    dp := make([]bool, n)
    res := 0
    for j := 0; j < n; j++ {
        for i := 0; i <= j; i++ {
            if i == j {
                res++
                dp[i] = true
            } else if j - i == 1 && s[i] == s[j] {
                dp[i] = true
                res++
            } else if j - i > 1 && s[i] == s[j] && dp[i+1] {
                dp[i] = true
                res++
            } else {
                dp[i] = false
            }
        }
    }
    return res
}
```


### 解码方法

``` javascript
func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }
    // 一维dp，优化空间复杂度
    cur, pre := 1, 1
    for i := 1; i < len(s); i++ {
        switch {
        // ‘10’， ‘20’这种情况，不会有额外的解码方法，保持不变，cur=pre
        case s[i] == '0' && (s[i - 1] == '1' || s[i - 1] == '2'):
            cur = pre
        // 除去上一种，其他的0都不合法，直接返回
        case s[i] == '0':
            return 0
        // 可以与前一个数字多一种解码方法的，加上之前的解码方法，并更新
        case (s[i] <= '6' && (s[i - 1] == '2' || s[i - 1] == '1')) || (s[i] > '6' && s[i - 1] == '1'):
            tmp := cur
            cur += pre
            pre = tmp
        // 没有多出解码方法的，cur保持不变， pre=cur
        default:
            pre = cur
        }
    }
    return cur
}
```

### 最大正方形


``` javascript
func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    res := 0
    // 初始化，并赋值，如果有1，结果为1，特殊处理
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                dp[i][j] = 1
                res = 1
            }
        }
    }
    // 0那一行，最多就是1，不需要进行更新，从1开始遍历
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 如果当前值是1,则以当前点为右下角包含1的正方形，为其他三个点的最小值+1
            // 为0的则 跳过
           if dp[i][j] == 1{
               dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
           }
           // 更新res
           if dp[i][j] > res {
               res = dp[i][j]
           }
        }
    }
    return res * res
}
```

### 最小路径和

``` javascript
func minPathSum(grid [][]int) int {
    // 初始化dp矩阵，与原数据矩阵对应
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    // 遍历矩阵，递推出dp[i][j]的值
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++  {
            switch {
            // 0,0点特殊处理
            case i == 0  && j == 0:
                dp[i][j] = grid[i][j]
            // 边界特殊处理，原矩阵当前点的 值加上之前的dp值，就是dp当前点的值
            case i == 0:
                dp[i][j] = dp[i][j - 1] + grid[i][j]
            case j == 0:
                dp[i][j] = dp[i - 1][j] + grid[i][j]
            // 非边界递推，最小路径，只能向右，向下，之前两个点的最小值，加上原矩阵当前点的值
            default:
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
            }
        }
    }
    return dp[m - 1][n - 1]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

```

### 矩形区域不超过k的最大和

``` javascript
// 前缀和加上最大子序和
func maxSumSubmatrix(matrix [][]int, k int) int {
    rowNum, colNum := len(matrix), len(matrix[0])
    // 最小值作为初始结果
    result := math.MinInt32
    // 按列遍历，起始指针
    for left := 0; left < colNum; left ++ {
        rowSum := make([]int, rowNum)
        // 按列遍历，结束指针
        for right := left; right < colNum; right++ {
            // 按行遍历，逐列想加，每一列都是之前列的行行和
            for row := 0; row < rowNum; row++ {
                rowSum[row] += matrix[row][right]
            }
            // 找最大值，每一列的最大子序和，就是矩形区域的和
            result = max(result, maxSubArrBelowK(rowSum, k))
            if result == k {
                return k
            }
        }
    }
    return result
}

// 找最大子序和，因为有k限制
func maxSubArrBelowK(arr []int, k int) int {
    // 先按动态规划找最大子序和，找到可以减少时间复杂度
    sum, max, l := arr[0], arr[0], len(arr)
    for i := 1; i < l; i++ {
        if sum > 0 {
            sum += arr[i]
        } else {
            sum = arr[i]
        }
        if sum > max {
            max = sum
        }
    }
    // 如果结果小于k，则找到，返回
    if max <= k {
        return max
    }
    // 结果大于k，则只能暴力法去找，最接近k的
    max = math.MinInt32
    for left := 0; left < l; left++ {
        sum := 0
        for right := left; right < l; right++ {
            sum += arr[right]
            if sum > max && sum <= k {
                max = sum
            }
            if max == k {
                return k
            }
        }
    }
    return max
}
```

### 编辑距离


``` javascript
func minDistance(word1 string, word2 string) int {
    m := len(word1)
    n := len(word2)
    // 初始化m+!,n+1,单词前面加一个空串处理
    dp := make([][]int, m+1)
    for i := 0; i < m + 1; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i < m+1; i++ {
        for j := 0; j < n+1; j++ {
            switch {
            // 00位置两个空串是0
            case i == 0 && j == 0:
                dp[i][j] = 0
            // 空串对于另一个单词就一直插入
            case i == 0:
                dp[i][j] = dp[i][j - 1] + 1
            case j == 0:
                dp[i][j] = dp[i - 1][j] + 1
            // 如果两个单词相等，则不需要操作
            case word1[i-1] == word2[j-1]:
                dp[i][j] = dp[i - 1][j - 1]
            // dp[i - 1][j - 1]代表替换，dp[i - 1][j]。i-1与j一样，删除i；i与j-1一样。i后面插入j
            default:
                dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
            }
        }
    }
    return dp[m][n]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}
```