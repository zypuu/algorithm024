// 有效字母异位词
package main

// 计数，最后看数字是否为0
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    countMap := make(map[byte]int, len(s))
    for i := 0; i < len(s); i++ {
        countMap[s[i]]++
        countMap[t[i]]--
    }
    for _, v := range countMap {
        if v != 0 {
            return false
        }
    }
    return true
}

func main() {
    
}