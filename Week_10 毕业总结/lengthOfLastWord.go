// 最后一个单词长度
package main

// 从后往前找第一个非空格，再往前找空格
func lengthOfLastWord(s string) int {
	tail := len(s) - 1
	for tail >= 0 && s[tail] == ' ' {
		tail--
	}
	if tail < 0 {
		return 0
	}
	head := tail
	for head >= 0 && s[head] != ' ' {
		head--
	}
	return tail - head
}
func main() {
    
}