// 省份数量
package main

// dfs
func findCircleNum(isConnected [][]int) int {
    var dfs func(i, j int)
    dfs = func(i,j int) {
        if isConnected[i][j] == 1 {
            isConnected[i][j] = 0
            isConnected[j][i] = 0
            for x := 0; x < len(isConnected); x++ {
                dfs(x, j)
            }
            for y := 0; y < len(isConnected[0]); y++ {
                dfs(i, y)
            }
        }
    }

    res := 0
    for i := 0; i < len(isConnected); i++ {
        for j := 0; j < len(isConnected[0]); j++ {
            if isConnected[i][j] == 1 {
                dfs(i, j)
                res++
            }
        }
    }
    return res
}

// 并查集1，每个元素单独处理，所有元素都加入集合
type UF struct {
    union []int
    count int
}

func newUf (m, n int, isConnected [][]int) *UF {
    uf := UF{make([]int, m * n), 0}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            uf.union[i * n + j] = i * n + j
            if isConnected[i][j] == 1 {
                uf.count++
            }
        }
    }
    return &uf
}

func(uf *UF) find(x int) int {
    root := x
    for uf.union[root] != root {
        root = uf.union[root]
    }
    for x != root {
        tmp := uf.union[x]
        uf.union[x] = root
        x = tmp
    }
    return root
}

func(uf *UF) Union(x, y int) {
    rootX := uf.find(x)
    rootY := uf.find(y)
    if rootX == rootY {
        return
    }
    uf.union[rootX] = rootY
    uf.count--
}

func findCircleNum(isConnected [][]int) int {
    m := len(isConnected)
    n := len(isConnected[0])
    uf := newUf(m, n, isConnected)
    // 看全部位置
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if isConnected[i][j] == 1 {
            	// 合并对称位置
                uf.Union(i*n + j, j*n + i)
                isConnected[i][j] = 0
                // 合并当前位置和对称位置的同一列
                for x := 0; x < m; x++ {
                    if isConnected[x][j] == 1 {
                        uf.Union(i*n + j, x*n + j)
                    }
                    if isConnected[x][i] == 1{
                        uf.Union(j*n + i, x*n + i)
                    }
                }
                // 合并当前位置和对称位置的同一行
                for y := 0; y < n; y++ {
                    if isConnected[i][y] == 1 {
                        uf.Union(i*n + j, i*n + y)
                    }
                    if isConnected[j][y] == 1 {
                        uf.Union(j*n + i, j*n + y)
                    }
                }
            }
        }
    }
    return uf.count
}

// 并查集2， 优化，初始化并查集直接按对称位置初始化
type UF struct {
    union []int
    count int
}

func newUf (m int) *UF {
    uf := UF{make([]int, m), m}
    for i := 0; i < m; i++ {
        uf.union[i] = i
    }
    return &uf
}

func(uf *UF) Union(x, y int) {
    rootX := uf.find(x)
    rootY := uf.find(y)
    if rootX == rootY {
        return
    }
    uf.union[rootX] = rootY
    uf.count--
}


func(uf *UF) find(x int) int {
    root := x
    for uf.union[root] != root {
        root = uf.union[root]
    }
    for x != root {
        tmp := uf.union[x]
        uf.union[x] = root
        x = tmp
    }
    return root
}

func findCircleNum(isConnected [][]int) int {
    m := len(isConnected)
    uf := newUf(m)
    // 只看对称位置
    for i := 0; i < m; i++ {
        for j := i; j < m; j++ {
            if isConnected[i][j] == 1 {
                uf.Union(i, j)
            }
        }
    }
    return uf.count
}

func main() {
    
}