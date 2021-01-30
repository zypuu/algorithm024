// 移动零
package main

import "fmt"

// moveZeroes ...
func moveZeroes(nums []int) []int  {
    i := 0
    for j:=0;j<len(nums);j++ {
        if nums[j] != 0 {
            // 交换位置
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
    return nums
}

func main() {
	res := moveZeroes([]int{2, 0, 6, 13, 0, 1})
	fmt.Print(res)
}