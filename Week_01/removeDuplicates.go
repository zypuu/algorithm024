// 删除数组重复项
package main

import "fmt"

// removeDuplicates ...
func removeDuplicates(nums []int) int {
    // 临界条件
    if len(nums) == 0 {
        return 0
    }
    a, b := 0, 1
    // 遍历
    for ;b < len(nums);b++ {
        if nums[b]!=nums[a] {
            nums[a+1] = nums[b]
            a++
        }
    }
    return a + 1
}

func main() {
	res := removeDuplicates([]int{1, 1, 2, 2, 2, 3})
	fmt.Print(res)
}