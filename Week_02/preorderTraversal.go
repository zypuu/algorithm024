// 二叉树的前序遍历
package main

// 根-左-右

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//preorderTraversal 前序遍历，递归方法，时间O(n),空间O(n)
func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    var handle func(*TreeNode)
    handle = func(root *TreeNode) {
        if root == nil {
            return
        }
        res = append(res, root.Val)
        handle(root.Left)
        handle(root.Right)
    }
    handle(root)
    return res
}

// preorderTraversal 迭代法，手动维护一个栈， 时间O(n),空间O(n)
func preorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    node := root
    res := []int{}

    for node != nil || len(stack) > 0 {
        for node != nil {
        	// 加入节点值
            res = append(res, node.Val)
            // 入栈
            stack = append(stack, node)
            // 根左右顺序
            node = node.Left
        }
        // 出栈，往右边走
        node = stack[len(stack) - 1].Right
        // 减小栈
        stack = stack[:len(stack) - 1]
    }
    return res
}


func main() {
	
}