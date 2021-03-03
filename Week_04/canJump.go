// 跳跃游戏
package main

// canJump ....
// 倒序贪心，从后开始，最先能跳到目标点的
func canJump(nums []int) bool {
    if len(nums) == 0 {
        return false
    }
    tmp := len(nums) - 1
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] + i >= tmp {
            tmp = i
        }
    }
    // 看最后是不是0
    return tmp == 0
}

func main() {
    
}