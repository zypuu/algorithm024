// 分发饼干
package main

// findContentChildren ....
// 先进行排序，然后贪心
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    i, j, res := 0, 0, 0
    // 通过两个指针，进行移动
    for i < len(g) && j < len(s) {
        if s[j] < g[i] {
            j++
        } else {
            res += 1
            i++
            j++
        }
    }
    return res
}

func main() {
    
}