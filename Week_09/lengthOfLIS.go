// 最长上升子序列
package main

// 动态规划
func lengthOfLIS(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    res := 1
    for i := 0; i < n; i++ {
        dp[i] = 1
        // 遍历当前点前面的最长上升子序列的dp值
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
            // 否则就是dp值就是1
        }
        res = max(dp[i], res)
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

// 二分加贪心
func lengthOfLIS(nums []int) int {
    res := 1
    n := len(nums)
    d := make([]int, n + 1)
    d[res] = nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] > d[res] {
            // 先++，再赋值
            res++
            d[res] = nums[i]
        } else {
            // 二分查找刚好大于nums[i]的
            s := 1
            e := res
            p := 0
            for s <= e {
                mid := s + (e - s) >> 1
                if nums[i] > d[mid] {
                    s = mid + 1
                    p = mid
                } else {
                    e = mid - 1
                }
            }
            // 找到下一位赋值，没找到p就是0，第一位赋值
            d[p + 1] = nums[i]
        }
    }
    return res
}

func main() {
    
}