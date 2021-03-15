// 最长有效括号
package main

// longestValidParentheses ....
// 最长 有效括号，双向遍历法
func longestValidParentheses(s string) int {
    if s == "" { return 0 }
    left, right, res := 0, 0, 0
    for i := 0; i < len(s); i++ {
        // 左右计数++
        if s[i] == '(' { left++ }
        if s[i] == ')' { right++ }
        // 右括号比左括号多，则重置，前面的舍弃
        if right > left { 
            left = 0
            right = 0
        }
        // 更新结果
        if left == right { res = max(res, right) }
    }
    // 以上会漏掉一种就是左括号始终比右括号多，这种是找不到的
    // 逆向遍历，正向排除右比左多的，逆向排除左比右多的
    left, right = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' { left++ }
        if s[i] == ')' { right++ }
        if right < left { 
            left = 0
            right = 0
        }
        if left == right { res = max(res, right) }
    }
    return 2 * res
}

// longestValidParentheses ...
// 动态规划方法
func longestValidParentheses(s string) int {
    if s == "" || len(s) < 2 {
        return 0
    }
    dp := make([]int, len(s))
    res := 0
    // 从1开始遍历
    for i := 1; i < len(s); i++ {
        // 结尾是）才能形成有效括号，只看结尾是）的前面， （的dp值都是0
        if s[i] == ')' {
            // 找前面的（，如果前一个就是（，前一个dp值是0，也能找到
            pre := i - 1- dp[i - 1]
            // 如果大于0，且找到是（，则结果+2
            if pre >= 0 && s[pre] == '(' {
                dp[i] = dp[i-1] + 2
                // 如果有再前一个，则加上再前一个的dp值
                if pre - 1 >= 0 { dp[i] += dp[pre - 1] }
            }
            // 返回最大值
            res = max(res, dp[i])
        }
        
    }
    return res 
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func main() {
    
}