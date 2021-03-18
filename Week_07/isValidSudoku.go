// 有效的数独
package main

// isValidSudoku
// 遍历通过map判断字符是否已经使用过
// 外层i，里层j， ij顺序代表一行一行遍历， ji顺序代表一列一列遍历
// 分块 行(i%3)*3 + j%3 列(i/3)*3 + j/3，得到的是一块9个的索引
func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        visr := make(map[byte]bool, 9)
        visc := make(map[byte]bool, 9)
        visb := make(map[byte]bool, 9)
        for j := 0; j < 9; j++ {
            // 行 ij
            if board[i][j] != '.' {
                if visr[board[i][j]] { return false }
                visr[board[i][j]] = true
            }
            // 列 ji
            if board[j][i] != '.' {
                if visc[board[j][i]] { return false }
                visc[board[j][i]] = true
            }
            // 分块
            row := (i%3)*3 + j%3
            col := (i/3)*3 + j/3
            if board[row][col] != '.' {
                if visb[board[row][col]] { return false }
                visb[board[row][col]] = true
            }
        }
    }
    return true
}

func main() {
    
}