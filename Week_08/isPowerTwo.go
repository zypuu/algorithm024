// 2的幂
package main

// 得到最低位的1，看是不是等于本身，是这说明就一个1
func isPowerOfTwo(n int) bool {
    if n == 0 { return false }
    return n & (-n) == n
}

// 清0最低位的1，看是不是等于0，是则说明只有1个1
func isPowerOfTwo(n int) bool {
    if n == 0 { return false }
    return n & (n-1) == 0
}

// 非位运算
func isPowerOfTwo(n int) bool {
    if n == 0 { return false }
    for  n % 2 == 0 {
        n /= 2
    }
    return n == 1
}

func main() {
    
}