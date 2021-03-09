// 最大正方形
package main

// maximalSquare ....
// 动态规划，dp[i][j]作为以当前点为右下角，所包含的正方形的最大边
func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    dp := make([][]int, m)
    res := 0
    // 初始化，并赋值，如果有1，结果为1，特殊处理
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                dp[i][j] = 1
                res = 1
            }
        }
    }
    // 0那一行，最多就是1，不需要进行更新，从1开始遍历
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // 如果当前值是1,则以当前点为右下角包含1的正方形，为其他三个点的最小值+1
            // 为0的则 跳过
           if dp[i][j] == 1{
               dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
           }
           // 更新res
           if dp[i][j] > res {
               res = dp[i][j]
           }
        }
    }
    return res * res
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}


func main() {
    
}