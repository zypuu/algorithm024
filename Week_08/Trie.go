// Trie树
package main

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

func main() {
    
}