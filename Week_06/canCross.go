// 青蛙过河
package main

// canCross ....
// dfs递归，剪枝
func canCross(stones []int) bool {
	// 记录在别的分支处理过的子问题
    vis := make(map[int]bool, len(stones))
    var dfs func(int, int) bool
    dfs = func(index, k int) bool {
    	// 如果在别的分支处理过，说明别的分支无法通过，能通过就说明不会有子问题，一直到底结束递归
    	// 如果找到，说明别的分支没到底，则1直接返回false
    	// 保证石头位置唯一
        if vis[index*100 + k] {
            return false
        }
        vis[index*100 + k] = true
        for i := index + 1; i < len(stones); i++ {
        	// 计算下一个石头与当前石头之间的距离
            dis := stones[i] - stones[index]
            // 在跳跃单位内，则继续递归，，return true
            if dis >= k - 1 && dis <= k + 1 {
                if dfs(i, dis) {
                    return true
                }
            // 超过了k+1，则说明无法到达，直接break
            } else if dis > k + 1 {
                break
            } // 小于k-1，说明太近了，可以下一波循环找后面的石头
        }
        // 判断当前分支最后是否到达最后位置
        return index == len(stones) - 1
    }
    return dfs(0, 0)
}

// 动态规划
// 动态规划
func canCross(stones []int) bool {
    // dp[stones[i]] 用个map保存第i个位置所有能跳的k值
    dp := make(map[int]map[int]int, len(stones))
    // 初始化每个石头，和能跳的k值
    for _, i := range stones {
        dp[i] = make(map[int]int)
    }
    // 第一个石头，0值，跳0步
    dp[0][0] = 0
    for i := 0; i < len(stones); i++ {
        for _, k := range dp[stones[i]] {
             // 初始0，-1， 0， 1，只有1步可以跳到下一个石头 
             for step := k - 1; step <= k + 1; step++ {
                //  如[0,1,3,5,6,8,12,17] 1 点能跳 0，1， 2 共3种距离 只有2跳到了 3里面  
                 if step > 0 {
                    // 跳的距离大于0 且 跳到数组中存在的位置 
                    if _, ok := dp[stones[i] + step]; ok {
                        // 则更新这个位置，dp【1】【0】 = 1
                        dp[stones[i] + step][i] = step
                    }           
                }
             }
        }
    }
    return  len(dp[stones[len(stones)-1]]) != 0
}

func main() {
    
}