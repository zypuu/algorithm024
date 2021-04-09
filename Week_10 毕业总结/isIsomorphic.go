// 同构字符串
package main

// 两个map，一次循环里，如果在map里看是不是等于t的
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) { return false }
    tmps := make(map[byte]byte, len(s))
    tmpt := make(map[byte]byte, len(t))
    for i := 0; i < len(s); i++ {
        if _, ok := tmps[s[i]]; !ok {
            tmps[s[i]] = t[i]
        }
        if _, ok := tmpt[t[i]]; !ok {
            tmpt[t[i]] = s[i]
        }
        if tmps[s[i]] != t[i] { return false }
        if tmpt[t[i]] != s[i] { return false }
        
    }
    return true
}

func main() {
    
}