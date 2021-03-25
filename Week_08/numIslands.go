// 岛屿数量
package main

// dfs
func numIslands(grid [][]byte) int {
    var dfs func(i,j int)
    dfs = func(i, j int) {
        if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && grid[i][j] == '1' {
            grid[i][j] = '0'
            dfs(i + 1, j)
            dfs(i - 1, j)
            dfs(i, j - 1)
            dfs(i, j + 1)
        }
    }
    
    res := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                dfs(i, j)
                res++
            }
        }
    }
    return res
}

// 并查集方法
// 并查集，记录count
type UF struct {
	union []int
    count int
}

// 初始化并查集
func newUF(m, n int, grid [][]byte) *UF {
	uf := UF{make([]int, m * n), 0}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
        	// 默认所有的1都是单独集合，索引位置是i*n+j
            uf.union[i * n + j] = i * n + j
            // 如果是1，集合数+1
            if grid[i][j] == '1' {
                uf.count++
            }
        }
    }
	return &uf
}
// 合并，合并成功，集合数-1
func (u *UF) Union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	u.union[rootX] = rootY
    u.count--
}
//  查找不需要修改
func (u *UF) find(x int) int {
	root := x
	for u.union[root] != root {
		root = u.union[root]
	}
	for x != root {
		tmp := u.union[x]
		u.union[x] = root
		x = tmp
	}
	return root
}
// 主函数
func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    uf := newUF(m, n, grid)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
        	// 如果是1，则看其他四个方向是不是1，是1就合并，然后置成0
            if grid[i][j] == '1' {
                if i - 1 >= 0 && grid[i-1][j] == '1' {
					uf.Union(i*n+j, (i-1)*n+j)
				}
				if i + 1 < m && grid[i+1][j] == '1' {
					uf.Union(i*n+j, (i+1)*n+j)
				}
				if j - 1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(i*n+j, i*n+(j-1))
				}
				if j + 1 < n && grid[i][j+1] == '1' {
					uf.Union(i*n+j, i*n+(j+1))
				}
				grid[i][j] = '0'
            }
        }
    }
    // 最后返回集合数即可
    return uf.count
}

func main() {
    
}