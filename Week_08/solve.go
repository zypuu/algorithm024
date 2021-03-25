// 被围绕的区域
package main

// dfs
func solve(board [][]byte)  {
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i >= 0 && j >= 0 && i < len(board) && j < len(board[0]) {
            if board[i][j] == 'O' {
            	// 先改成别的，最后再遍历统一更新
                board[i][j] = 'Y'
                dfs(i - 1, j)
                dfs(i + 1, j)
                dfs(i, j - 1)
                dfs(i, j + 1)
            } 
        } 
    }

    for i := 0; i < len(board); i++ {
        if board[i][0] == 'O' { dfs(i, 0) }
        if board[i][len(board[0]) - 1] == 'O' { dfs(i, len(board[0]) - 1) }
    }
    for j := 0; j < len(board[0]); j++ {
        if board[0][j] == 'O' { dfs(0, j) }
        if board[len(board) - 1][j] == 'O' { dfs(len(board) - 1, j) }
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] != 'Y' {
                board[i][j] = 'X'
            } else {
                board[i][j] = 'O'
            }
        }
    }
}

// 并查集
type UF struct {
    parents []int
}

func newUF (n int) *UF {
    uf := UF{make([]int, n)}
    for i := 0; i < n; i++ {
        uf.parents[i] = i
    }
    return &uf
}

func(uf *UF)find(x int) int {
    root := x
    for root != uf.parents[root] {
        root = uf.parents[root]
    }
    for x != root {
        tmp := uf.parents[x]
        uf.parents[x] = root
        x = tmp
    }
    return root
}

func(uf *UF)union(x, y int) {
    rootX := uf.find(x)
    rootY := uf.find(y)
    if rootX == rootY {
        return
    }
    uf.parents[rootX] = rootY
}

func(uf *UF)isConnect(x, y int) bool {
    return uf.find(x) == uf.find(y)
}

func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    t := m * n
    // 多创建一个集合，边界元素放到该集合内，以该集合为准
    uf := newUF(t + 1)
  
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                // 边界位置都归到一个单独的集合
                if i == 0 || j == 0 || i == m - 1|| j == n -1 {
                    uf.union(i * n + j, t)
                }
                // 只判断左边和上边，即可覆盖全部元素
                if i > 0 && board[i - 1][j] == 'O' { uf.union(i *n + j, (i - 1)*n + j) }
                if j > 0 && board[i][j - 1] == 'O' { uf.union(i *n + j, i*n + j - 1) }
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // 判断是O的如果不和边界的元素属于同一集合，，则是被包围，更新为X
            if board[i][j] == 'O' && !uf.isConnect(i * n + j, t) {
                board[i][j] = 'X'
            }
        }
    }
}

func main() {
    
}