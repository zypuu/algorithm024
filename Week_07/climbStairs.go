// 爬楼梯
package main

var vis = map[int]int{1: 1, 2: 2}

// climbStairs ....
// 递归 法，记忆化递归，记录计算过的值
func climbStairs(n int) int {
    if _, ok := vis[n]; ok {
        return vis[n]
    }
    vis[n] = climbStairs(n - 1) + climbStairs(n - 2)
    return vis[n]
}

func main() {
    
}