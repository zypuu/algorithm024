// 第一个唯一字符
package main

// map统计，返回等于1的
func firstUniqChar(s string) int {
    if s == "" {
        return -1
    }
    countMap := make(map[byte]int, len(s))
    for i := 0; i < len(s); i++ {
        countMap[s[i]]++
    }
    for i := 0; i < len(s); i++ {
        if countMap[s[i]] == 1 {
            return i
        }
    }
    return -1
}

func main() {
    
}