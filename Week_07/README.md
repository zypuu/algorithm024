学习笔记

DFS模板
``` javascript
// 深度优先，栈遍历
func dfs(root *TreeNode) {
	// root不存在则返回
	if root == nil {
		return
	}
	var dfs func()
	dfs = func() {
		if visit[root] {
			return
		}
		visit[root] = true
		for i := range root.Children {
			if !visit[i] {
				dfs(i)
			}
			
		}
	}
	
}

```

BFS模板
``` javascript
// 广度优先，队列遍历
func bfs(root *TreeNode) {
	// root不存在则返回
	if root == nil {
		return
	}
	// 手动维护一个队列，初始化，并放入root，还有一个visited的map
	q := []*TreeNode{root}
	visited := map[*TreeNode]bool{}
 	// 未访问则记录,处理当前层
 	for len(q) > 0 {
 		// 节点出队
 		node := q[0]
 		q = q[1:]
 		// 访问记录
 		visited[node] = true
 		// 处理逻辑
 		process(node)
 		// 更新q，更新node
 		nodes = related(node)
 		q = append(q, nodes...)
 	}
	
}

```

双向BFS模板

``` javascript
// 双向bfs
func tebfs(root *TreeNode) {
	// 开始队列
	startQ := []interface{}
	// 结束队列
	endQ := []interface{}
	// 开始使用过的记录
	visitedStart := map[int]bool
	// 结束使用记录
	visitedEnd := map[int]bool
	// 全局使用记录
	visited := map[int]bool
	
	// 遍历
	for len(startQ) > 0 {
		// 遍历开始队列
		qLen := len(startQ)
		for i := 0; i < qLen; i++ {
			cur := q[0]
			q = q[1:]
			// 处理逻辑
			process()
			// 开始记录处理
			visitedStart()
			// 放入队列
			startQ = append(startQ, ...)
		}
		// 层遍历完处理逻辑,互换队列，互换记录map
		if len(startQ) > len(endQ) {
			startQ, endQ = endQ, startQ
			visitedStart, visitedEnd = visitedEnd, visitedStart
		}
	}
}
```
A* 模板
``` javascript
func ASearch() {
	// 初始化优先队列
	pq = priorityQ()

	pq =append(pq, start) // 获取优先队列

	visited := map[int]bool{start : true}
	for len(pq) > 0 {
		node := pq.get() // 根据优先级获取
		visited[node] = true

		process()
		// 添加进队列
		pq = append(node.next)
	} 	
}
```

