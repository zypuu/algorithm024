// 不同路径2
package main

// uniquePathsWithObstacles 
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])

    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }
    // 注意起始点障碍物
    if obstacleGrid[0][0] == 1{
        dp[0][0] = 1
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // 有障碍就是0，并跳过
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
                continue
            }
            switch {
            case i == 0 && j == 0: dp[i][j] = 1
            case i == 0: dp[i][j] = dp[i][j - 1]
            case j == 0: dp[i][j] = dp[i - 1][j]
            default: dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
            }
        }
    }
    return dp[m - 1][n - 1]
}

func main() {
    
}