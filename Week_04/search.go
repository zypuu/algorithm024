// 搜索旋转数组
package main

// search ...
// 二分查找，不同情况比对，特殊条件处理缩小的条件
func search(nums []int, target int) int {
    // 边界条件
    if len(nums) == 1 {
        if target == nums[0] {
            return 0
        } else {
            return -1
        }
    }
    // 左右下标
    l := 0
    r := len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        // 等于目标值返回索引
        if nums[mid] == target {
            return mid
        }
        // 分四种情况
        // 左边小于中间，说明左边升序，旋转在右边
        if nums[l] <= nums[mid]  {
            // target在左边，则往左边缩小
            if target >= nums[l] && target <= nums[mid] {
                r = mid - 1
            } else { // 否则右边缩小
                l = mid + 1
            }
        // 否则右边升序
        } else {
            // target在右边，往右边缩小，否则左边
            if target >= nums[mid] && target <= nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1 
}

// 如果数组包含重复数字，搜索旋转数组二
func search(nums []int, target int) bool {
    if len(nums) == 1 {
        if target == nums[0] {
            return true
        }
        return false
    }
    l := 0
    r := len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] == target {
            return true
        }
        // 对重复数字进行处理跳过
        if nums[l] == nums[mid] {
            l++
        } else if nums[mid] == nums[r] {
            r--
        } else if nums[l] < nums[mid] {
            if target >= nums[l] && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if target > nums[mid] && target <= nums[r] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return false
}

func main() {
    
}