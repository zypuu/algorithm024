// 接雨水
package main

import "fmt"


/*
思路：涉及雨水计算，要有3个索引

1.从0开始遍历，第一个值无法计算面积，直接入栈，cur++，高度为1

2.到第二个值，cur为1，高度为0，小于之前入栈的高度为1，继续入栈

3.到第三个值，cur为2，栈不为0，高度为2，大于栈里面存储的之前cur的高度，也就是0，开始计算雨水，说明中间是凹下去的
   则得到栈顶的高度为0，也就是第二个索引的高度
   栈出掉第二个索引
   计算距离，当前索引-第一个索引-1，也就是中间夹的雨水宽
   找到两边高度最小的
   最小高度减去中间凹下去的高度
   计算雨水面积


4.然后cur+1入栈，继续下一个索引，之前计算过的已经出栈

5.重复以上操作，每次计算都是当前高度与下一个比它高的高度之间的雨水面积

*/


// trap 接雨水 stack栈思想
func trap(height []int) int {
    sum := 0
    // 初始化栈，存放高度的索引值
    stack := make([]int, 0)
    // 当前遍历值初始化
    current := 0
    // for
    for current < len(height) {
        // 如果栈不为0，且 当前的高度大于栈顶的高度，计算雨水面积一定涉及3个索引
        for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
            // 当前高度为第3个索引，h为第二个索引
            h := height[stack[len(stack)-1]]
            // 去除最后一位，栈顶元素为第一个索引
            stack = stack[0 : len(stack)-1]
            // for循环终止，栈为0
            if len(stack) == 0 {
                break
            }
            // 当前索引值-栈顶的第一个的索引-1，求宽
            distance := current - stack[len(stack)-1] - 1
            // 找到当前索引高度与第一个索引低的那个高度
            min := min(height[stack[len(stack)-1]], height[current])
            // 计算雨水，低的那个高度，减去第二个索引的高度
            sum += distance * (min - h)
        }
        // 当前索引值入栈
        stack = append(stack, current)
        // 当前索引加一
        current++

    }
    // 最后返回结果
    return sum
}

// min ...
func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}


func main() {
	res := trap([]int{1,0,2})
	fmt.Print(res)
}