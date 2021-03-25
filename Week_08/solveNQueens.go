 // N皇后
package main

// 基于数组的递归
func solveNQueens(n int) [][]string {
    res := [][]string{}
    tmp := make([][]string, n)
    for i := 0; i < n; i++ {
        tmp[i] = make([]string, n)
        for j := 0; j < n; j++ {
            tmp[i][j] = "."
        }
    }
    visc := make(map[int]bool, n)
    visr := make(map[int]bool, n)
    visb := make(map[int]bool, n)

    var dfs func(r int)
    dfs = func(r int) {
        if r == n {
            cur := make([]string, 0, n)
            for i := 0; i < n; i++ {
                cur =append(cur, strings.Join(tmp[i], ""))
            }
            res = append(res, cur)
            return
        }

        for i := 0; i < n; i++ {
            if visc[i] || visr[r - i] || visb[r + i] {
                continue
            }
            tmp[r][i] = "Q"
            visc[i] = true
            visr[r - i] = true
            visb[r + i] = true
            dfs(r + 1)
            visc[i] = false
            visr[r - i] = false
            visb[r + i] = false
            tmp[r][i] = "."
        }
    }
    dfs(0)
    return res
}

// 基于位运算的递归
func solveNQueens(n int) [][]string {
	// 最终结果
    res := [][]string{}
    // 初始化一个slice，记录每一行q的位置，初始化都为-1
    bSlice := make([]int, n)
    for i := 0; i < n; i++ {
        bSlice[i] = -1
    }

    var dfs func(int, int, int, int)
    // 递归函数
    // row行。 col列。 dia1.dia2两个对角线
    dfs = func(row, col, dia1, dia2 int) {
        // 递归终止条件,符合规则的一种结果
        if row == n {
            tmp := []string{}
            // 遍历每一行，每一列，初始化都为 .
            for i := 0; i < n; i++ {
                cur := make([]byte, n)
                for j := 0; j < n; j++ {
                    cur[j] = '.'
                }
                // 根据每一行的q的位置，更新q
                cur[bSlice[i]] = 'Q'
                // byte数组转为string加入结果
                tmp = append(tmp, string(cur))
            }
            res = append(res, tmp)
            return
        }
        // 列，两个对角线，或运算(col | dia1 | dia2)，即所有使用了的位置
        // ^对以上结果取反，得到还未使用的位置，即可用的位置
        // (1 << n) - 1)从最高位到第n位，都清0，得到n位二进制
        availP := ((1 << n) - 1) & (^(col | dia1 | dia2))
        // 循环递归，可用位置不为0
        for availP != 0 {
        	// 取最低位的可用位置
            p := availP & (-availP)
            // 最低位清0，更新可用位置
            availP &= availP - 1
            // 减去1之后计算1的个数，得到这一行的q的索引位置，也就是p的1在第几位
            // 高位->低位就是bslice从左->右
            bSlice[row] = bits.OnesCount(uint(p - 1))
            // 继续递归。行数+1， 列与最低位的可用位置取或，代表这一位已经使用了
            // 对角线与可用位置取或，再根据对角线方向，一个左移，一个右移
            dfs(row + 1, col | p, (dia1 | p) >> 1, (dia2 | p) << 1)
            // 回溯，这一行的q的位置为-1
            bSlice[row] = -1
        }
    }

    dfs(0, 0, 0, 0)
    return res
}

func main() {
    
}