// 比特位计数
package main

// 非dp
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++ {
        cur := i
        for cur != 0 {
            cur &= cur - 1
            res[i]++
        }
    }
    return res
}

// 高有效位
func countBits(num int) []int {
    dp := make([]int, num+1)
    hb := 0
    for i := 1; i <= num; i++ {
        // 如果是2的幂，则更新高有效位 
        if i & (i - 1) == 0 { 
            hb = i
        }
        // 高有效位的下一位是+1，以此类推，直到下一个2的幂，并更新
        dp[i] = dp[i - hb] + 1
    }
    return dp
}

// 奇偶数，等于去掉最后一位前面的值 加上新的一位，是奇数就是1，偶数就是0
func countBits(num int) []int {
    dp := make([]int, num+1)
    for i := 1; i <= num; i++ {
        dp[i] = dp[i >> 1] + i & 1
    }
    return dp
}

func main() {
    
}