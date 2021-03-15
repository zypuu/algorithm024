// 矩形区域不超过K的最大和
package main

// maxSumSubmatrix ....
// 前缀和加上最大子序和
func maxSumSubmatrix(matrix [][]int, k int) int {
    rowNum, colNum := len(matrix), len(matrix[0])
    // 最小值作为初始结果
    result := math.MinInt32
    // 按列遍历，起始指针
    for left := 0; left < colNum; left ++ {
        rowSum := make([]int, rowNum)
        // 按列遍历，结束指针
        for right := left; right < colNum; right++ {
            // 按行遍历，逐列想加，每一列都是之前列的行行和
            for row := 0; row < rowNum; row++ {
                rowSum[row] += matrix[row][right]
            }
            // 找最大值，每一列的最大子序和，就是矩形区域的和
            result = max(result, maxSubArrBelowK(rowSum, k))
            if result == k {
                return k
            }
        }
    }
    return result
}

// 找最大子序和，因为有k限制
func maxSubArrBelowK(arr []int, k int) int {
    // 先按动态规划找最大子序和，找到可以减少时间复杂度
    sum, max, l := arr[0], arr[0], len(arr)
    for i := 1; i < l; i++ {
        if sum > 0 {
            sum += arr[i]
        } else {
            sum = arr[i]
        }
        if sum > max {
            max = sum
        }
    }
    // 如果结果小于k，则找到，返回
    if max <= k {
        return max
    }
    // 结果大于k，则只能暴力法去找，最接近k的
    max = math.MinInt32
    for left := 0; left < l; left++ {
        sum := 0
        for right := left; right < l; right++ {
            sum += arr[right]
            if sum > max && sum <= k {
                max = sum
            }
            if max == k {
                return k
            }
        }
    }
    return max
}

func max(num1, num2 int) int {
    if num1 > num2 {
        return num1
    }
    return num2
}

func main() {
    
}