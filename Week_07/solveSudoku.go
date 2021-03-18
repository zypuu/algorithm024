// 解数独
package main

// solveSudoku
// 递归回溯
func solveSudoku(board [][]byte)  {
	// 初始化，通过每一行的索引数字判断（每一列同理）
    var visr, visc [9][9]bool
    // 初始化块的，块的索引，以及里面的数字的索引判断
    var visb [3][3][9]bool
    // 初始化需要填入的点
    var spaces [][2]int
    // 预处理数据
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
        	// 需要填入的点加入列表
            if board[i][j] == '.' {
                spaces = append(spaces, [2]int{i, j})
            // 不需要填入的点，根据位置设置true
            } else {
            	// byte转int，-1
                b := board[i][j] - '1'
                visr[i][b] = true
                visc[j][b] = true
                visb[i/3][j/3][b] = true
            }
        }
    }
    // 递归函数
    var dfs func(n int) bool
    dfs= func(n int) bool {
    	// 遍历到最后一个点结束后，终止
        if n == len(spaces) {
            return true
        }
        // 取要填入的点
        i, j := spaces[n][0], spaces[n][1]
        for k := byte(0); k < 9; k++ {
        	// 判断是否已经使用过，有则跳过
            if visr[i][k] || visc[j][k] || visb[i/3][j/3][k] {
                continue
            }
            // 没有记录
            board[i][j] = k + '1'
            visr[i][k] = true
            visc[j][k] = true
            visb[i/3][j/3][k] = true
            // 递归下一个点，如果下一个点返回true，则返回true
            if dfs(n + 1) {
                return true
            }
            // 回溯
            board[i][j] = '.'
            visr[i][k] = false
            visc[j][k] = false
            visb[i/3][j/3][k] = false
        }
        // 最终如果都不返回true，则返回false，失败
        return false
    }
    dfs(0)
}

func main() {
    
}