// 回文串
package main

// countSubstrings ....
func countSubstrings(s string) int {
    n := len(s)
    dp := make([]bool, n)
    res := 0
    for j := 0; j < n; j++ {
        for i := 0; i <= j; i++ {
        	// 一个字符
            if i == j {
                res++
                dp[i] = true
            // 两个字符且相等
            } else if j-i == 1 && s[i] == s[j] {
                dp[i] = true
                res++
            // 两个以上，且相等，且里面的是回文
            } else if j-i > 1 && s[i] == s[j] && dp[i+1] {
                dp[i] = true
                res++
            } else {
            	// 不符合改回false，下一个外层从头更新
                dp[i] = false
            }
        }
    }
    return res
}

func main() {
    
}