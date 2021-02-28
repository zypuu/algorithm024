// 岛屿数量
package main

// numIslands 
func numIslands(grid [][]byte) int {
    res := 0
    var dfs func(int, int)
    dfs = func(x, y int) {
        // 相邻的都置成0
        if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) &&  grid[x][y] == '1' {
            grid[x][y] = '0'
            dfs(x + 1, y)
            dfs(x - 1, y)
            dfs(x, y + 1)
            dfs(x, y - 1)
        }
    }
    // 遍历二维数组，发现是1，则是岛屿，把与它相邻的1都置0，都是相连的属于一个岛屿
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                res += 1
                dfs(i, j)
            }
        }
    }
    return res
}