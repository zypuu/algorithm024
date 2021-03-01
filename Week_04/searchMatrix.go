// 搜索二维矩阵数组
package main

// searchMatrix
// 缩减搜索范围，先找到对应的行。然后再对该行进行二分
func searchMatrix(matrix [][]int, target int) bool {
    cur := []int{}
    for _, v := range matrix {
        if target >= v[0] && target <= v[len(v) - 1] {
            cur = v
        }
    }
    if len(cur) == 0 {
        return false
    }
    l := 0
    r := len(cur) - 1
    for l <= r {
        mid := (l + r) / 2
        if cur[mid] == target {
            return true
        }
        if cur[mid] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false
}

// 直接二分法
// 将二维数组想象成一个完整的一维数组，通过一维数字的索引，转换成二维数组的索引
// 通过一维数组的索引进行二分，每一行的个数是n，整数结果是第几行，取余是第几列
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    l := 0
    r := m * n - 1
    for l <= r {
        mid := l + (r - l) / 2
        if matrix[mid / n][mid % n] == target {
            return true
        }
        if matrix[mid / n][mid % n] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return false
}