// 最长有效括号
package main

// 动态规划
func longestValidParentheses(s string) int {
    if s == "" || len(s) < 2 {
        return 0
    }
    dp := make([]int, len(s))
    res := 0
    for i := 1; i < len(s); i++ {
        // 只看）的
        if s[i] == ')' {
            // 找到前面的，且减去前面的有效括号的长度，长度如果增加，则包在里面的也一定是有效括号
            pre := i - 1 - dp[i - 1]
            if pre >= 0 && s[pre] == '(' {
                dp[i] = dp[i - 1] + 2
                if pre - 1 >= 0 {
                    dp[i] += dp[pre - 1]
                }
            }
            if dp[i] > res {
                res = dp[i]
            }
        }
    }
    return res
}

// 双向遍历法
func longestValidParentheses(s string) int {
    if s == "" || len(s) < 2 {
        return 0
    }
    res := 0
    left, right := 0, 0
    // 从前往后，左=右说明形成一个有效括号，找最长，但是会漏掉一种情况，所以需要反向再来一遍
    for i := 0; i < len(s); i++ {
        if s[i] == '(' { left++ }
        if s[i] == ')' { right++ }
        if right > left { left, right = 0, 0 }
        if left == right {
            if right > res { res = right }
        }
    }
    // 从右往左
    left, right = 0, 0
    for i := len(s) - 1; i >= 0 ; i-- {
        if s[i] == '(' { left++ }
        if s[i] == ')' { right++ }
        if left > right { left, right = 0, 0 }
        if left == right {
            if right > res { res = right }
        }
    }
    return res * 2
}

func main() {
    
}