// 组合
package main

// combine 递归方法
// 1.递归中往数组里添加数字，不断往深递归，达到k的长度要求，则加入结果
// 2.达到结果开始回溯，当前结果满足，则回到上一递归结果，回到上一递归状态
// 3.继续上一个递归的for循环，i++，继续开始下一个数字的组合
func combine(n int, k int) [][]int {
    res := [][]int{}
    if n < k {
        return res
    }
    var handle func(start int, path []int)
    handle = func(start int, path []int) {
        // 得到结果的长度，加入结果，加入后进行回溯
        if len(path) == k {
        	// 后续会修改path，需要拷贝，不然数据会修改，path是地址引用
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        } 
        for i := start; i <= n; i++ {
            // 1加入
            path = append(path, i)
            // 从2开始，判断是否到回溯条件，如果符合，加入最终结果
            handle(i+1, path)
            // 回溯去掉最后一位，继续递归。2完了3，开始下一位循环
            path = path[:len(path) - 1]
        }

    }
    handle(1, []int{})
    return res
}

func main() {
    
}