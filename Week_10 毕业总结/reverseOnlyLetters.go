// 仅仅反转字母
package main

// 两个指针跳过，交换位置
func reverseOnlyLetters(S string) string {
    tmp := []byte(S)
    for i, j := 0, len(tmp)-1; i < j; {
        for i < j && !isLetter(tmp[i]) {
            i++
        }
        for i < j && !isLetter(tmp[j]) {
            j--
        }
        tmp[i], tmp[j] = tmp[j], tmp[i]
        i++
        j--
    }
    return string(tmp)
}

func isLetter(c byte) bool {
    if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
        return true
    }
    return false
}

func main() {
    
}