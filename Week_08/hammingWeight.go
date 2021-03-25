// 1的个位，汉明重量
package main

// 清0次数
func hammingWeight(num uint32) int {
    count := 0
    for num > 0 {
        // 将最低位的1清0
        num &= num - 1
        count++
    }
    return count
}

// 循环检查
func hammingWeight(num uint32) int {
    count := 0
    for i := 0; i < 32; i++ {
        if (num >> i) & 1 == 1{
            count++
        }
    }
    return count
}

func main() {
    
}