// 反转字符串2
package main

// 双指针，每次加2k
func reverseStr(s string, k int) string {
    tmp := []byte(s)
    i := 0
    for i < len(tmp) {
        l := i
        r := i + k - 1
        if r > len(s) - 1{
            r = len(s) - 1
        }
        for k := l; k < r; k++ {
            tmp[k], tmp[r] = tmp[r], tmp[k]
            r--
        }
        i += 2 * k
    }
    return string(tmp)
}

func main() {
    
}