// 有效字母的异位词
package main

import "fmt"
import "sort"
import "strings"

// isAnagram1 排序暴力法 时间O(nlogn)，空间O(n)
func isAnagram1(s string, t string) bool {
    a := strings.Split(s, "")
    b := strings.Split(t, "")
    sort.Strings(a)
    sort.Strings(b)
    s = strings.Join(a, "")
    t = strings.Join(b, "")
    return s == t
}

// isAnagram2 map法,同时解决unicode字符问题,时间O(n)，空间O(n)
func isAnagram2(s string, t string) bool {
    n := make(map[rune]int, len(s))
    for i, v := range s + t {
        if i < len(s) {
            n[v] += 1
        } else {
            n[v] -= 1
        }
    }
    for k, _ := range n {
        if n[k] != 0 {
            return false
        }
    }
    return true
}

func main() {
	res := isAnagram2("abcd", "dcba")
	fmt.Print(res)
}