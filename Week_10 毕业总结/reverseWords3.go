// 反转字符串里的单词3
package main

// 空格分隔，反转每一个
func reverseWords(s string) string {
    tmp := strings.Split(s, " ")
    for i := 0; i < len(tmp); i++ {
        tmp[i] = reverSingle([]byte(tmp[i]))
    }
    return strings.Join(tmp, " ")
}

func reverSingle(s []byte) string {
    for i := 0; i < len(s) >>1; i++ {
        s[i], s[len(s) - 1- i] = s[len(s) - 1- i], s[i]
    }
    return string(s)
}

func main() {
    
}