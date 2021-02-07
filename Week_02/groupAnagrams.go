// 异位词分组
package main

import "fmt"
import "sort"
import "strings"

// groupAnagrams1 map法 质数做key 时间O(n(k+26))，空间O(n(k+26))
func groupAnagrams1(strs []string) [][]string {
    var res [][]string
    resMap := make(map[int][]string, len(strs))
    prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    for _, v := range strs {
        num := 1
        for _, i := range v {
        	// 质数相乘保证唯一
            num *= prime[int(rune(i)) - 97]
        }
        if _, ok := resMap[num]; ok {
            resMap[num] = append(resMap[num], v)
        } else {
            resMap[num] = []string{v}
        }
    }
    // 最后遍历添加结果
    for _, v := range resMap {
        res = append(res, v)
    }
    return res
}

func main() {
	res := groupAnagrams1(["abcd", "dcba"])
	fmt.Print(res)
}