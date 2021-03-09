// 解码方法
package main

// numDecodings ....
// 动态规划，dp[i]表示当前字符串的解码方法
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

func main() {
    
}