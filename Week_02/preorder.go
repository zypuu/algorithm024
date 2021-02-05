// N叉树的前序遍历
package main

import "fmt"



type Node struct {
    Val int
    Children []*Node
}
 
var res []int

//preorder 递归方法
func preorder(root *Node) []int {
    // 重置
    res = []int{}
    handle(root)
    return res
}

// 递归函数
func handle(root *Node) {
    if root != nil {
        res = append(res, root.Val)
        // 子节点
        for _, v := range root.Children {
            handle(v)
        }
    }
}

//preorder 迭代法
func preorder(root *Node) []int {
    var res []int
    // 辅助栈
    var stack = []*Node{root}
    for 0 < len(stack) {
        // 节点遍历，for更新root
        for root != nil {
            //前序输出，根左右，将根节点加入结果
            res = append(res, root.Val)
            // 子节点为0，就跳出循环
            if 0 == len(root.Children) {
                break
            }
            // 遍历子节点，子节点入栈，当做下次循环的根节点，倒序加入，不加入第一个，栈先入后出
            for i := len(root.Children) - 1; 0 < i; i-- {
                stack = append(stack, root.Children[i]) //入栈
            }
            // 更新root，为子节点的第一个，根->左
            root = root.Children[0]
        }
        // 没有子节点则跳出循环，出栈，最底层的右，根->左->右
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
    }
    return res
}


func main() {
	res = preorder(&Node{Val: 1, Children: []&Node{&Node{Val: 3, Children: []&Node{nil}}}})
	fmt.Print(res)
}