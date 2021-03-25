// 颠倒二进制
package main

func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i := 0; i < 32; i++ {
    	// 新数字左移1位
        res <<= 1
        // num & 1是获取num的最后一位，也可以判断奇偶
        // 最后一位加1，用或也可以，因为res最后一位肯定是0
        res |= num & 1
        // 右移1位，更新num
        num >>= 1
    }
    return res
}

// 16位互相交换
func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i := 0; i < 16; i++ {
        res |= (num & (1 << i)) << (31 - 2 * i) 
    }
    for i := 16; i < 32; i++ {
        res |= (num & (1 << i)) >> (2 * i -31)
    }
    return res
}

func main() {
    
}