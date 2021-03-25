### 二叉树的层序遍历

``` javascript
// dfs
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, l int) {
        if l == len(res) {
            res = append(res, []int{})
        }
        res[l] = append(res[l], root.Val)
        if root.Left != nil { dfs(root.Left, l + 1) }
        if root.Right != nil { dfs(root.Right, l + 1) }
    }
    dfs(root, 0)
    return res
}

// bfs
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    q := []*TreeNode{root}

    for len(q) > 0 {
        qLen := len(q)
        curRes := []int{}
        for i := 0; i < qLen; i++ {
            node := q[0]
            q = q[1:]
            curRes = append(curRes, node.Val)
            if node.Left != nil { q = append(q, node.Left) }
            if node.Right != nil { q = append(q, node.Right) }
        }
        res = append(res, curRes)
    }
    return res
}
```

注意判断边界条件，开始的root是否为nil， root的left，right是否为nil


### trie 字典树


``` javascript
type Trie struct {
    isEnd       bool
    children    [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    root := this
    // 每个单词字符处理，是nil的话就新建
    for i := 0; i < len(word); i++ {
        if root.children[word[i] - 'a'] == nil {
            root.children[word[i] - 'a'] = &Trie{}
        }
        // 不是nil就更新下一节点去遍历
        root = root.children[word[i] - 'a']
    }
    // 最后一个节点是单词结尾
    root.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    root := this
    // 搜索与插入类似，是nil的话直接返回false
    for i := 0; i < len(word); i++ {
        if root.children[word[i] - 'a'] == nil {
            return false
        }
        root = root.children[word[i] - 'a']
    }
    // 最后判断是前缀，还是完整单词
    if root.isEnd != true {
        return false
    }
    return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    root := this
    // 只判断前缀，不需要判断完整单词
    for i := 0; i < len(prefix); i++ {
        if root.children[prefix[i] - 'a'] == nil {
            return false
        }
        root = root.children[prefix[i] - 'a']
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
```

### 单词搜索

``` javascript
// dfs
func exist(board [][]byte, word string) bool {
    dirI := []int{-1, 1, 0, 0}
    dirJ := []int{0, 0, 1, -1}
    var dfs func(int, int, int) bool
    dfs = func(i, j, l int) bool {
        if board[i][j] != word[l] {
            return false
        }
        if l == len(word) - 1 {
            return true
        }
        board[i][j] = ' '
        defer func(i, j int, b byte)  { board[i][j] = b }(i, j, word[l])
        for k := 0; k < 4; k++ {
            newI := i + dirI[k]
            newJ := j + dirJ[k]
            if newI >= 0 && newJ >= 0 && newI < len(board) && newJ < len(board[0]) && board[newI][newJ] != ' ' {
                if dfs(newI, newJ, l + 1) { return true }
            }
        }
        return false
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] {
                if dfs(i, j, 0) { return true }
            }
        }
    }
    return false
}

// Trie 字典树， 注释看单词搜索2，一样
type Trie struct {
    isEnd bool
    childrens [58]*Trie
}

func exist(board [][]byte, word string) bool {
    root := &Trie{}
    node := root
    for _, v := range word {
        if node.childrens[v - 'A'] == nil {
            node.childrens[v - 'A'] = &Trie{}
        }
        node = node.childrens[v - 'A']
    }
    node.isEnd = true

    dirx := []int{-1, 0, 1, 0}
    diry := []int{0, 1, 0, -1}

    var dfs func(x, y int, root *Trie) bool
    dfs = func(x, y int, root *Trie) bool {
        if x < 0 || y < 0 || x == len(board) || y == len(board[0]) {
            return false
        }
        cur := board[x][y]
        if cur == ' ' || root.childrens[cur - 'A'] == nil {
            return false
        }
        root = root.childrens[cur - 'A']
        if root.isEnd {
            return true
        }
        board[x][y] = ' '
        for p := 0; p < 4; p++ {
            if dfs(x + dirx[p], y + diry[p], root) {
                return true
            }
        }
        board[x][y] = cur
        return false
    }
    for x := 0; x < len(board); x++ {
        for y := 0; y < len(board[0]); y++ {
            if dfs(x, y, root) {
                return true
            }
        }
    }
    return false
}
```

### 单词搜索II

``` javascript
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
```

单词搜索II，字典树方法的时间复杂度
O（m * n * 4^L）m,n是二维网格的长宽，L是单词的平均长度


### 并查集

代码模板


``` javascript
type UF struct {
	union []int
}

// 初始化
func newUF(cap int) *UF {
	uf := UF{
		make([]int, cap),
	}
	// 每个元素的索引等于自己,每一个元素是一个集合，索引也代表元素
	// 每个元素是代表元素
	for i := 0; i < cap; i++ {
		uf.union[i] = i
	}
	return &uf
}

// 合并
func (u *UF) Union(x, y int) {
	// 查找两个元素的代表元素
	rootX := u.find(x)
	rootY := u.find(y)
	// 如果代表元素一样 ，则属于同一集合
	if rootX == rootY {
		return
	}
	// 如果不是则x元素的代表元素位置的代表元素改为y的代表元素
	u.union[rootX] = rootY
}
// 判断两个元素是不是同一集合。也就是两个位置上的root代表元素是不是同一个
func (u *UF) Connected(x, y int) bool {
	return u.find(x) == u.find(y)
}

// 查询一个元素的代表元素，属于哪个集合
func (u *UF) find(x int) int {
	root := x
	// x元素位置是不是x自己，如果不是，说明该元素被合并过，有别的代表元素
	// 被合并过后x位置上的元素就是它的父节点
	// 此过程是不断往上找 父节点，找他的代表元素
	for u.union[root] != root {
		root = u.union[root]
	}
	// 压缩路径，将该条分支上的所有子节点的位置上的元素都是root
	// 先取x位置上的节点，然后x位置上改为root
	// 再把x位置改为刚取出的节点的位置，看是不是root
	for x != root {
		tmp := u.union[x]
		u.union[x] = root
		x = tmp
	}
	// 最后返回root节点，也就是x的代表元素
	return root
}

```

### 岛屿数量

``` javascript
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
```

### 被围绕区域

``` javascript
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
```

### 省份数量

``` javascript
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

```

### 位运算

#### 与 &

有两个数都是1结果才为1

``` javascript
	var b uint8 = 20 // 00010100
	var c uint8 = 15 // 00001111
	a := b & c // 00000100
	a = 4
```

#### 或 |

两个数有一个是1 结果就是1

``` javascript
	var b uint8 = 20 // 00010100
	var c uint8 = 15 // 00001111
	a := b & c // 00011111
	a = 31
```

#### 异或 ^

 ^作二元运算符就是异或，相同为0，不相同为1

``` javascript
	var b uint8 = 20 // 00010100
	var c uint8 = 15 // 00001111
	a := b & c // 00011011
	a = 27
```

^作一元运算符表示是按位取反

``` javascript
	var b uint8 = 20 // 00010100
	var c uint8 = 15 // 00001111
	a := b & c // 00011011
	^a = 228 // 11100100
```

uint8与int
``` javascript
	var b uint8 = 20 // 00010100
	var c int = 20 // 00010100
	^b = 228 // 11101011
	^c = -21 // -10101
```

特殊情况
``` javascript
	var b uint8 = 20 // 00010100
	b ^ b = 0
	b ^ 0 = b
```


int类型，故最高位是符号位，符号位取反，所以得到的结果是负数

一个有符号位的^操作为 这个数+1的相反数

#### &^

将运算符左边数据相异的位保留，相同位清零

1&^1  得0
1&^0  得1
0&^1  得0
0&^0  得0

``` javascript
	var b uint8 = 20 // 00010100
	var c uint8 = 15 // 00001111
	a := b &^ c
	a = 16 // 00010000
```
以b为准，相异保留b的位，相同，清零b的位

#### << 左移 >> 右移

左移规则：右边空出的位用0填补，高位左移溢出则舍弃该高位

右移规则：左边空出的位用0或者1填补。正数用0填补，负数用1填补。注：不同的环境填补方式可能不同。低位右移溢出则舍弃该位

``` javascript
	var b uint8 = 20 // 00010100
	a := b << 1
	a = 40 // 00101000
	c：= b >> 1
	c = 10 // 00001010
```
左移1相当于乘以2
右移1相当于除以2

#### 判断奇偶

同时也是获取最后一位的值

``` javascript
	var b int = 22 // 10110
	a := b & 1
	a == 0 // 偶数
	a == 1 // 奇数
```

#### 清0最低位的1
``` javascript
	var b int = 22 // 10110
	a := b & (b - 1)
	a = 20 // 10100
```

#### 得到最低位的1

``` javascript
	var b int = 22 // 10110
	a := b & (-b)
	a = 2 // 10
```

#### 将第n+1位置为1

``` javascript
	var b int = 22 // 10110
	a := b | (1 << 0) // 将第0位置为1,1 << n, 第n位
	a = 23 // 10111
```

#### 将第n+1为置为0

``` javascript
	var b int = 22 // 10110
	a := b & (^(1 << 1)) // 将第1位置为0,1 << n, 第n位
	a = 20 // 10100
```

#### 将右n位清0

``` javascript
	var b int = 22 // 10110
	a := b & (^0 << 2)// 将右边2位清0， n，右边n位
	a = 20 // 10100
```

#### 获取第n+1位值

``` javascript
	var b int = 22 // 10110
	a := (b >> 2) & 1 // 获取第2位值
	a = 1 // 1
```

#### 获取第n+1位的幂值

``` javascript
	var b int = 22 // 10110
	a := b & (1 << 2) // n=2，其实是从右边数第3位的
	a = 4 // 从右边开始算，第一位是0， n
```

####  将最高位至第n+1位（含）清0

``` javascript
	var b int = 22 // 10110
	a := b & ((1 << 3) - 1) // n=3，其实是从左边数到第3位，不含第3位，前面的都清0
	a = 6 // 110
```

### 颠倒二进制


``` javascript
func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i := 0; i < 32; i++ {
    	// 新数字左移1位
        res <<= 1
        // num & 1是获取num的最后一位，也可以判断奇偶
        // 最后一位加1，用或也可以，因为res最后一位肯定是0
        res |= num & 1
        // 右移1位，更新num
        num >>= 1
    }
    return res
}

// 16位互相交换
func reverseBits(num uint32) uint32 {
    var res uint32 = 0
    for i := 0; i < 16; i++ {
        res |= (num & (1 << i)) << (31 - 2 * i) 
    }
    for i := 16; i < 32; i++ {
        res |= (num & (1 << i)) >> (2 * i -31)
    }
    return res
}
```

### 1的个位

``` javascript
// 清0次数
func hammingWeight(num uint32) int {
    count := 0
    for num > 0 {
    	// 将最低位的1清0
        num &= num - 1
        count++
    }
    return count
}

// 循环检查
func hammingWeight(num uint32) int {
    count := 0
    for i := 0; i < 32; i++ {
        if (num >> i) & 1 == 1{
            count++
        }
    }
    return count
}
```

### 2的幂

``` javascript
// 得到最低位的1，看是不是等于本身，是这说明就一个1
func isPowerOfTwo(n int) bool {
    if n == 0 { return false }
    return n & (-n) == n
}

// 清0最低位的1，看是不是等于0，是则说明只有1个1
func isPowerOfTwo(n int) bool {
    if n == 0 { return false }
    return n & (n-1) == 0
}

// 非位运算
func isPowerOfTwo(n int) bool {
    if n == 0 { return false }
    for  n % 2 == 0 {
        n /= 2
    }
    return n == 1
}
```

### 比特计数

``` javascript
// 非dp算法
func countBits(num int) []int {
    res := make([]int, num+1)
    for i := 1; i <= num; i++ {
        cur := i
        for cur != 0 {
            cur &= cur - 1
            res[i]++
        }
    }
    return res
}

// 高有效位
func countBits(num int) []int {
    dp := make([]int, num+1)
    hb := 0
    for i := 1; i <= num; i++ {
        // 如果是2的幂，则更新高有效位 
        if i & (i - 1) == 0 { 
            hb = i
        }
        // 高有效位的下一位是+1，以此类推，直到下一个2的幂，并更新
        dp[i] = dp[i - hb] + 1
    }
    return dp
}

// 奇偶数，等于去掉最后一位前面的值 加上新的一位，是奇数就是1，偶数就是0
func countBits(num int) []int {
    dp := make([]int, num+1)
    for i := 1; i <= num; i++ {
        dp[i] = dp[i >> 1] + i & 1
    }
    return dp
}
```

### N皇后

基于位运算的回溯递归

``` javascript
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
```