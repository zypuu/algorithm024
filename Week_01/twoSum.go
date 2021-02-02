// 两数之和
package main

import "fmt"

// twoSum map方法
func twoSum(nums []int, target int) []int {
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

func main() {
	res := twoSum([]int{2,7,11,15}, 9)
	fmt.Print(res)
}