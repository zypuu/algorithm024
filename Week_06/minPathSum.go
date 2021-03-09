// 最小路径和
package main

// minPathSum ....
// 动态规划，dp[i][j]作为到达当前点的最小路径和，结果就是dp[m - 1][n - 1]
func minPathSum(grid [][]int) int {
    // 初始化dp矩阵，与原数据矩阵对应
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    // 遍历矩阵，递推出dp[i][j]的值
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++  {
            switch {
            // 0,0点特殊处理
            case i == 0  && j == 0:
                dp[i][j] = grid[i][j]
            // 边界特殊处理，原矩阵当前点的 值加上之前的dp值，就是dp当前点的值
            case i == 0:
                dp[i][j] = dp[i][j - 1] + grid[i][j]
            case j == 0:
                dp[i][j] = dp[i - 1][j] + grid[i][j]
            // 非边界递推，最小路径，只能向右，向下，之前两个点的最小值，加上原矩阵当前点的值
            default:
                dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
            }
        }
    }
    return dp[m - 1][n - 1]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}


func main() {
    
}