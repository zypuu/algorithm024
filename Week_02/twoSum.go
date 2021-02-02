// 两数之和
package main

import "fmt"

// twoSum1 map方法
func twoSum1(nums []int, target int) []int {
    resMap := make(map[int]int, len(nums))
    for i, v := range nums {
        // 判断另一个数在不在map，在则找到返回结果
        if j, ok := resMap[target - v]; ok {
            return []int{j, i}
        }
        // 不在则把该数添加进map
        resMap[v] = i
    }
    return []int{}
}

// twoSum2 暴力解法，注意嵌套for循环的写法，避免重复计算
func twoSum2(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{j, i}
            }
        }
    }
    return []int{}
}

func main() {
	res := twoSum1([]int{2,7,11,15}, 9)
	fmt.Print(res)
}