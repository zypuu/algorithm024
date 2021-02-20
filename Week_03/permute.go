// 全排列
package main

// permute 递归法，分治回溯
// 1.递归中往数组里添加数字，不断往深递归，达到结果要求，加入最终结果
// 2.满足一种结果后，进行回溯，全排列，相当于每一个空位，每一个数字都有可能
// 3.避免重复，对于已经加入的数字通过map记录，因为是全遍历，再加入时进行过滤
// 4.回溯到上一状态，重置之前加入的数字，重置当前数组
func permute(nums []int) [][]int {
    // 结果
    res := [][]int{}
    // 用于记录访问过的节点
    vis := make(map[int]bool, len(nums))
    var handle func(curNums []int)
    handle = func(curNums []int) {
        // 找到结果的条件
        if len(curNums) == len(nums) {
            // 数组是地址引用，需要拷贝加入最终结果
            tmp := make([]int, len(curNums))
            copy(tmp, curNums)
            res = append(res, tmp)
            return
        }
        // 数组全遍历，要对重复数字进行过滤
        for _, v := range nums {
            if vis[v] {
                continue
            }
            // 加入结果，并记录
            vis[v] = true
            curNums = append(curNums, v)
            handle(curNums)
            // 回溯，重置记录与数组
            vis[v] = false
            curNums = curNums[:len(curNums) - 1]
        }
    }
    handle([]int{})
    return res
}

func main() {
    
}