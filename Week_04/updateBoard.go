// 扫雷游戏
package main

// 深度优先，递归
func updateBoard(board [][]byte, click []int) [][]byte {
    // 8个方向的使用
    var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
    var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

    // 如果点击的是M，则是游戏结束，更新值
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
        return board
    }
    // 参数是坐标值，和每次递归更新后的board
    var dfs func([][]byte, int, int)
    dfs = func(board [][]byte, x, y int) {
        // 记录周围的M数量
        mCount := 0
        for i := 0; i < 8; i++ {
            tx := x + dirX[i]
            ty := y + dirY[i]
            // 边界条件
            if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) { continue }
            // 如果周围有M，则数字+1
            if board[tx][ty] == 'M' {
                mCount++
            }
        }
        // 周围有M，则更新为byte数字
        if mCount > 0 {
            board[x][y] = byte(mCount + '0')
        } else {
            // 周围没有M则更新为B
            board[x][y] = 'B'
            // 向8个方向扩展
            for i := 0; i < 8; i++ {
                tx := x + dirX[i]
                ty := y + dirY[i]
                // 如果周边的不是E，说明已经被访问过了，则跳过
                if tx < 0 || ty < 0 || tx >= len(board) || ty >= len(board[0]) || board[tx][ty] != 'E' { continue }
                dfs(board, tx, ty)
            }
        }        
    }
    dfs(board, click[0], click[1])
    return board
}


// 广度优先，bfs
func updateBoard(board [][]byte, click []int) [][]byte {
    // 8个方向
    var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
    var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

    // 点击到M结束
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
        return board
    }
    var bfs func([][]byte, int, int)
    bfs = func(board [][]byte, x, y int) {
        q := [][]int{}
        vis := make([][]bool, len(board))
        // 初始化所有的节点，是否被访问过
        for i := 0; i < len(vis); i++ {
            vis[i] = make([]bool, len(board[0]))
        }
        // 当前点击加入队列
        q = append(q, []int{x, y})
        // 记录当前点击访问
        vis[x][y] = true
        // 遍历队列
        for i := 0; i < len(q); i++ {
            // 获取当前坐标，初始化周围M的数量
            mCount, cx, cy := 0, q[i][0], q[i][1]
            // 找周围的M
            for i := 0; i < 8; i++ {
                tx, ty := cx + dirX[i], cy + dirY[i]
                if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
                    continue
                }
                // 记录周围的M
                if board[tx][ty] == 'M' {
                    mCount++
                }
            }
            // 如果有M，更新值
            if mCount > 0 {
                board[cx][cy] = byte(mCount + '0')
            } else {
                // 没有M更新为B
                board[cx][cy] = 'B'
                // 将周围的加入队列
                for i := 0; i < 8; i++ {
                    tx, ty := cx + dirX[i], cy + dirY[i]
                    // 这里不需要在存在 B 的时候继续扩展，因为 B 之前被点击的时候已经被扩展过了
                    if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' || vis[tx][ty] {
                        continue
                    }
                    q = append(q, []int{tx, ty})
                    vis[tx][ty] = true
                }
            }
    }
    }
    bfs(board, click[0], click[1])
    

    return board
}
