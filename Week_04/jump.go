// 跳跃游戏2
package main

// jump ...
// 最少的跳跃次数，则从前开始贪心，每次跳最大步数
// 也就是从后开始，可以到达最后结果的最远步数，下标最小的那个
func jump(nums []int) int {
    count := 0
    tmp := len(nums) - 1
    for tmp > 0 {
        // 每次从头开始循环，找到最先满足的值，然后break，更新tmp
        // 每次都是最先满足，最优解，贪心算法
        for i := 0; i < tmp; i++ {
            if nums[i] + i >= tmp {
                count += 1
                tmp = i
                break
            }
        }
    }
    return count
}

// 优化贪心，减少for循环
// 从前往后找其每次能跳跃的最大步数，实现贪心
func jump(nums []int) int {
    count := 0
    end := 0
    maxTmp := 0
    // 一次循环
    for i := 0; i < len(nums) - 1; i++ {
        // 每次循环找其跳的最大步数
        maxTmp = max(maxTmp, nums[i] + i)
        // 遍历到end，说明要发生跳跃，加1，并更新end
        if i == end {
            count += 1
            end = maxTmp
        }
    }
    return count
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}