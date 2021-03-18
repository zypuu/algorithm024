// N皇后
package main

// solveNQueens
// 递归回溯剪枝
func solveNQueens(n int) [][]string {
    res := [][]string{}
    // 初始化结果
    bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	// 列，左右对角线过滤剪枝
    visCol := make(map[int]bool, n)
    visLeftSth := make(map[int]bool, n)
    visRightSth := make(map[int]bool, n)
    // 传入行数
    var handle func(int)
    handle = func(r int) {
    	// 行数到n则加入结果
        if r == n {
            tmp := make([]string, n)
            for i := 0; i < n; i++ {
                tmp[i] = strings.Join(bd[i], "")
            }
            res = append(res, tmp)
            return
        }
        // 行内遍历列
        for c := 0; c < n; c++ {
        	// 剪枝判断列，左右对角线有没有皇后
            if visCol[c] || visLeftSth[c - r] || visRightSth[c + r] {
                continue
            }
            // 没有改为Q，并记录
            bd[r][c] = "Q"
            visCol[c] = true
            visLeftSth[c - r] = true
            visRightSth[c + r] = true
            // 行数加1，继续递归
            handle(r + 1)
            // 回溯
            bd[r][c] = "."
            visCol[c] = false
            visLeftSth[c - r] = false
            visRightSth[c + r] = false
        } 
    }
    handle(0)
    return res
}

func main() {
    
}