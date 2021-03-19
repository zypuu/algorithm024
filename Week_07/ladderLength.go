// 单词接龙
package main

// ladderLength 
// 单向bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
    vis := make(map[string]bool, len(wordList))
    for i := 0; i < len(wordList); i++ {
        vis[wordList[i]] = true
    }
    // 初始化队列
    q := []string{beginWord}
    // count是for之后++
    count := 1
    for len(q) > 0 {
    	// 记录当前队列长度，遍历当前队列，q长度会变
        qLen := len(q)
        for k := 0; k < qLen; k++ {
        	// 出队
            cur := q[0]
            q = q[1:]
            if cur == endWord {
                return count
            }
            for i := 0; i < len(cur); i++ {
                for j := 'a'; j <= 'z'; j++ {
                	// 得到新单词
                    newWord := cur[:i] + string(j) + cur[i+1:]
                    // 在map里
                    if vis[newWord] {
                        q = append(q, newWord)
                        vis[newWord] = false
                    }
                }
            }
        }
        count++ 
    }
    return 0
}

// 双向bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 初始化startq， endq
    startQ := []string{beginWord}
    endQ := []string{endWord}
    // 初始化从前从后的两个map，记录是否使用过
    visStart := make(map[string]bool, len(wordList))
    visEnd := make(map[string]bool, len(wordList))
    // end是最终结果
    visEnd[endWord] = true
    // 单词map，用于判断是不是在库里
    vis := make(map[string]bool, len(wordList))
    for _, v := range wordList {
        vis[v] = true
    }
    // 如果结尾不在库里，返回0
    if _, ok := vis[endWord]; !ok {
        return 0
    }
    // count是for之后++
    count := 1
    // 遍历队列
    for len(startQ) > 0 {
        qLen := len(startQ)
        for i := 0; i < qLen; i++ {
        	// 出队
            cur := startQ[0]
            startQ = startQ[1:]
            // 拼新单词
            for c := 0; c < len(cur); c++ {
                for j := 'a'; j <= 'z'; j++ {
                    new := cur[:c] + string(j) + cur[c+1:]
                    // 如果不在库里，或者起始的记录过则跳过
                    if !vis[new] || visStart[new] {
                        continue
                    }
                    // 如果在结束的map里，这说明匹配到了
                    if visEnd[new] {
                        return count + 1
                    }
                    // 加入队列，记录start
                    startQ = append(startQ, new)
                    visStart[new] = true
                }
            }
        }
        // 层级遍历完++
        count++
        // 互换队列，保证每次遍历的队列长度最小，互换记录
        // 每次遍历都是 startq
        if len(startQ) > len(endQ) {
            startQ, endQ = endQ, startQ
            visStart, visEnd = visEnd, visStart
        }

    }
    return 0
}


func main() {
    
}