// 翻转对
package main

// 归并排序解法
func reversePairs(nums []int) int {
    if len(nums) == 0 { return 0 }
    count := 0
    // 开始排序
    mergeSort(nums, &count, 0, len(nums) - 1)
    return count
}

func mergeSort(nums []int, count *int, start, end int) {
    // 终止条件
    if start >= end { return }
    mid := start + (end - start) >> 1
    mergeSort(nums, count, start, mid)
    mergeSort(nums, count, mid + 1, end)

    // 下面是针对排序好的数组去计算
    i := start
    j := mid + 1
    // 左边的升序，右边的升序
    for i <= mid && j <= end {
        // 统计翻转对
        if nums[i] > 2 * nums[j] {
            *count += mid - i + 1
            j++
        } else {
            i++
        }
    }
    // 还原，合并数组，注意申请的数组下标
    i = start
    j = mid + 1
    tmp := make([]int, 0, end - start + 1)
    for i <= mid && j <= end {
        if nums[i] < nums[j] {
            tmp = append(tmp, nums[i])
            i++
        } else {
            tmp = append(tmp, nums[j])
            j++
        }
    }
    // 没合并的左边数组
    if i <= mid {
        tmp = append(tmp, nums[i:mid + 1]...)
    }
    // 没合并的右边数组
    if j <= end {
        tmp = append(tmp, nums[j:end + 1]...)
    }
    // 赋值nums，更新数组
    k := 0
    for i := start; i <= end; i++ {
        nums[i] = tmp[k]
        k++
    }
}

func main() {
    
}