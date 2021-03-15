// 编辑举例
package main

// minDistance ....
// 动态规划，dp[i][j]作为两个单词到i的位置。与j位置的单词所需要的最小编辑次数
func minDistance(word1 string, word2 string) int {
    m := len(word1)
    n := len(word2)
    // 初始化m+!,n+1,单词前面加一个空串处理
    dp := make([][]int, m+1)
    for i := 0; i < m + 1; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i < m+1; i++ {
        for j := 0; j < n+1; j++ {
            switch {
            // 00位置两个空串是0
            case i == 0 && j == 0:
                dp[i][j] = 0
            // 空串对于另一个单词就一直插入
            case i == 0:
                dp[i][j] = dp[i][j - 1] + 1
            case j == 0:
                dp[i][j] = dp[i - 1][j] + 1
            // 如果两个单词相等，则不需要操作
            case word1[i-1] == word2[j-1]:
                dp[i][j] = dp[i - 1][j - 1]
            // dp[i - 1][j - 1]代表替换，dp[i - 1][j]。i-1与j一样，删除i；i与j-1一样。i后面插入j
            default:
                dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1
            }
        }
    }
    return dp[m][n]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}


func main() {
    
}