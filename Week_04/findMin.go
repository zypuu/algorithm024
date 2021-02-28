// 搜索旋转数组最小值，也是转折点
package main

// findMin
// 二分查找 
func findMin(nums []int) int {
    if len(nums) < 2 {
        return  nums[0]
    }
    l := 0
    r := len(nums) - 1
    for l <= r {
        // 最后到同一位就是最小值，也就是转折点
        if l == r {
            return nums[l]
        }
        mid := (l + r) / 2
        // 说明后半是有转折点的区间，否则就在前半
        if nums[r] < nums[mid] {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return 0
}

// findMin 库函数，相比于二分都是logN，但是空间复杂度有额外的数组
func findMin(nums []int) int {
    sort.Ints(nums)
    return nums[0]
}