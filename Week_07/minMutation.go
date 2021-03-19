// 最小基因变化
package main

// minMutation ...
// 单向bfs
func minMutation(start string, end string, bank []string) int {
    vis := make(map[string]bool, len(bank))
    // 判断是否重复，避免死循环
    for i := 0; i < len(bank); i++ {
        vis[bank[i]] = true
    }
    q := []string{start}
    item := []string{"A", "C", "G", "T"}
    // 这是转换次数，单词接龙是单词数目，注意
    res := 0
    for len(q) > 0 {
        qLen := len(q)
        for i := 0; i < qLen; i++ {
        	// 出队
            cur := q[0]
            q = q[1:]
            // 找到结果
            if cur == end {
                return res
            }
            // 循环每个字符
            for i := 0; i < len(cur); i++ {
                for j := 0; j < 4; j++ {
                    new := cur[:i] + item[j] + cur[i+1:]
                    if vis[new] {
                        q = append(q, new)
                        vis[new] = false
                    }
                }
            }
        }
        // 遍历完一层，结果++
        res++
    }
    return -1
}

// 双向bfs
func minMutation(start string, end string, bank []string) int {
	// 初始化
	startQ := []string{start}
    endQ := []string{end}

    visStart := make(map[string]bool, len(bank))
    visEnd := make(map[string]bool, len(bank))
    visEnd[end] = true

    // 判断end在不在库，初始化库
    vis := make(map[string]bool, len(bank))
    for _, v := range bank {
        vis[v] = true
    }
    if _, ok := vis[end];!ok {
        return -1
    }
    // 可更改字符
    item := []string{"A", "C", "G", "T"}
    res := 0
    for len(startQ) > 0 {
        qLen := len(startQ)
        for i := 0; i < qLen; i++ {
            cur := startQ[0]
            startQ = startQ[1:]
            for c := 0; c < len(cur); c++ {
                for _, v := range item {
                    new := cur[:c] + v + cur[c+1:]
                    // 不在库里，或者已使用过则跳过
                    if !vis[new] || visStart[new] {
                        continue
                    }
                    // 在结束里，返回结果+1
                    if visEnd[new] {
                        return res + 1
                    }
                    startQ = append(startQ, new)
                    visStart[new] = true
                }
            }
        }
        res++
        // 互换队列
        if len(startQ) > len(endQ) {
            startQ, endQ = endQ, startQ
            visStart, visEnd = visEnd, visStart
        }
    }
    return -1
}

func main() {
    
}