// 同构字符串
package main

// 两个map互相判断
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    tmpS := make(map[byte]byte, len(s))
    tmpT := make(map[byte]byte, len(s))
    for i := 0; i < len(s); i++ {
        if _, ok := tmpS[s[i]]; !ok {
            tmpS[s[i]] = t[i]
        }
        if _, ok := tmpT[t[i]]; !ok {
            tmpT[t[i]] = s[i]
        }
        if tmpS[s[i]] != t[i] { return false }
        if tmpT[t[i]] != s[i] { return false }
    }
    return true
}

func main() {
    
}