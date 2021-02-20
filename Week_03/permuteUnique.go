// 全排列2，包含重复数字
package main

// permuteUnique 全排列2，包含重复数字
// 1.递归不断往数组中添加数字，往深递归，达到结果长度，加入最终结果
// 2.之后回溯，重置状态
// 3.相比全排列1，包含了重复数字，除了要对已经加入的进行记录过滤，还要对重复数字进行过滤
// 4.为了更好的处理重复，对数组进行排序，重复数字都在一起
// 5.因为包含重复数字，所以记录下标，判断下标是否被记录
// 6.、已经记录的过滤，当前下标值与之前的一样，遇到重复的，且前面一个没有被记录的话，说明是一个新的循环，会产生重复结果，要跳过
func permuteUnique(nums []int) [][]int {
    res := [][]int{}
    vis := make(map[int]bool, len(nums))
    // 数组排序，方便处理重复数据
    sort.Ints(nums)

    var handle func(curNums []int)
    handle = func(curNums []int) {
        if len(curNums) == len(nums) {
            tmp := make([]int, len(curNums))
            copy(tmp, curNums)
            res = append(res, tmp)
            return
        }
        for i, v := range nums {
            // 已经记录过的跳过
            if vis[i] {
                continue
            }
            // 与前面的数字值重复，且前面的数字没有被记录过，说明是新的循环，说明会产生重复结果，跳过
            if i > 0 && v == nums[i - 1] && !vis[i - 1] {
                continue
            }
            // 加入数字，记录
            vis[i] = true
            curNums = append(curNums, v)
            handle(curNums)
            // 回溯重置
            vis[i] = false
            curNums = curNums[:len(curNums) - 1]
        }
    }

    handle([]int{})
    return res
}

func main() {
    
}