// 解码方法
package main

// 动态规划，不同情况判断
func numDecodings(s string) int {
    cur, pre := 0, 1
    if s[0] != '0' {
        cur = 1
    }
    for i := 1; i < len(s); i++ {
        switch {
        case s[i] == '0' && (s[i - 1] == '1' || s[i - 1] == '2'):
            cur = pre
        case s[i] == '0':
            return 0
        case s[i - 1] == '1' || (s[i - 1] == '2' && s[i] <= '6'):
            tmp := cur
            cur += pre
            pre = tmp
        default:
            pre = cur
        }
    }
    return cur
}

func main() {
    
}