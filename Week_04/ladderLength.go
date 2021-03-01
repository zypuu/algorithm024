// 单词接龙
package main

// ladderLength
// 单词接龙，通过bfs，维护一个队列
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 用来记录，单词是否被使用过，使用过的单词改为false
	wordMap := map[string]bool{}
	for _,w := range wordList{
		wordMap[w] = true
	}
	// 初始化队列
	queue := []string{beginWord}
	// 初始化次数，默认肯定有一次
	level := 1
	// 遍历队列
	for len(queue) != 0 {
		// 记录长度，因为遍历过程中会修改队列
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			// 单词出队
			word := queue[0]
			// 队列减少1
			queue = queue[1:]
			// 如果该单词是结束单词，则返回次数
			if word == endWord {
				return level
			}
			// 然后找到该单词的可替换单词，26个字母遍历
			for c := 0; c < len(word); c++ {
				for j := 'a'; j <= 'z'; j++ {
					// 得到新单词
					newWord := word[:c] + string(j) + word[c+1:]
					// 如果得到的新单词在map里，说明满足条件，入队，并将该单词置为false
					// 如果有多个词，则队列里有多个，队列里是可能满足当前跳转的单词
					if wordMap[newWord] == true { 
						queue = append(queue, newWord) 
						wordMap[newWord] = false    
					}
				}
			}
		}
		// 当前levelSize遍历完之后，完成接龙一次，计数
		level++
	}
	return 0
}
