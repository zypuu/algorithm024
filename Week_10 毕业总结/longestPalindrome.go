// 最长回文串
package main

// 中心扩展法
func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    start, end := 0, 0
    // 遍历，中心扩展
    for i := 0; i < len(s); i++ {
        // 奇数情况
        l1, r1 := expandS(s, i, i)
        // 偶数情况
        l2, r2 := expandS(s, i, i + 1)
        if r1 - l1 > end - start {
            start = l1
            end = r1
        }
        if r2 - l2 > end - start {
            start = l2
            end = r2
        }
    }
    return s[start:end + 1]
}

func expandS(s string, l, r int) (int, int) {
    for l >= 0 && r <= len(s) - 1 && s[l] == s[r] {
        l--
        r++
    }
    return l + 1, r - 1
}

// 动态规划
func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    start := 0
    maxLen := 1
    for i := 0; i < len(s); i++ {
        for j := 0; j < i; j++ {
            switch {
            // 同一个字符串，肯定是
            case i == j: dp[j][i] = true
            // 相等的情况下，如果长度小于2，肯定是，或者里面的是回文串
            case s[i] == s[j] && (i - j <= 2 || dp[j + 1][i - 1]): dp[j][i] = true
            default:  dp[j][i] = false 
            }
            // 记录最大长度与起始位置
            if dp[j][i] {
                curLen := i - j + 1
                if curLen > maxLen {
                    maxLen = curLen
                    start = j
                }
            }
        }
    }
    return s[start:start + maxLen]
}

func main() {
    
}