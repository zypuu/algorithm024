// 单词搜索2
package main

// trie树方法
// 创建字典树
type Trie struct {
    word string
    childrens [26]*Trie
}
// 方向
var dirx = []int{-1, 1, 0, 0}
var diry = []int{0, 0, -1, 1}
// 主函数
func findWords(board [][]byte, words []string) []string {
    res := []string{}
    // 初始化root
    root := &Trie{}
    for _, v := range words {
    	// 每个单词重新从root开始
        node := root
        // 插入字典树
        for i := 0; i < len(v); i++ {
            if node.childrens[v[i] - 'a'] == nil {
                node.childrens[v[i] - 'a'] = &Trie{}
            }
            node = node.childrens[v[i] - 'a']
        }
        node.word = v
    }
    // 递归
    var dfs func(x, y int, root *Trie)
    dfs = func(x, y int, root *Trie) {
    	// 超出边界，return
        if x < 0 || y < 0 || x == len(board) || y == len(board[0]) {
            return
        }
        // 取当前点
        cur := board[x][y]
        // 如果当前点使用过，或者不在字典树里，则return
        if cur == ' ' || root.childrens[cur - 'a'] == nil {
            return
        }
        // 在字典树里，且没有使用过，则更新到下一个节点
        root = root.childrens[cur - 'a']
        // 如果到结束，找到单词，加入结果
        if root.word != "" {
            res = append(res, root.word)
            root.word = ""
        }
        // 标记防止重复
        board[x][y] = ' '
        // 继续dfs
        for p := 0; p < 4; p++ {
            dfs(x + dirx[p], y + diry[p], root)
        }
        // 还原
        board[x][y] = cur
    }
    // 遍历 dfs
    for x := 0; x < len(board); x++ {
        for y := 0; y < len(board[0]); y++ {
            dfs(x, y, root)
        }
    }
    return res
}


//dfs
func findWords(board [][]byte, words []string) []string {
    m := len(board)
    n := len(board[0])
    res := make([]string, 0, len(words))
    for i := 0; i < len(words); i++ {
        flag := false
        for j := 0; j < m; j++ {
            for k := 0; k < n; k++ {
                flag = dfs(board, words, j, k, i, 0)
                if flag { break }
            }
            if flag { break }
        }
        if flag { res = append(res, words[i]) }
    }
    return res
}

var dirx = []int{-1, 1, 0, 0}
var diry = []int{0, 0 ,-1, 1}

func dfs(board [][]byte, words []string, j, k, i, t int) bool {
    if board[j][k] != words[i][t] {
        return false
    }
    if t == len(words[i]) - 1 {
        return true
    }
    board[j][k] = ' '
    // 还原board
    defer func(j, k int, b byte) { board[j][k] = b } (j, k, words[i][t])
    for d := 0; d < 4; d++ {
        newJ := j + dirx[d]
        newK := k + diry[d]
        if newJ >= 0 && newK >= 0 && newJ < len(board) && newK < len(board[0]) && board[newJ][newK] != ' ' {
            if dfs(board, words, newJ, newK, i, t + 1) { return true }
        }
    }
    return false
}

func main() {
    
}