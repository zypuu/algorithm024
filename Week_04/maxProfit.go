// 买卖股票的最佳时机
package main

// maxProfit 贪心算法
// 每次循环，只要有收益就进行计算，完成收益累加
func maxProfit(prices []int) int {
    res := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i - 1] {
            res += prices[i] - prices[i - 1]
        }
    }
    return res
}

func main() {
    
}