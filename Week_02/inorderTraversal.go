// 二叉树的中序遍历
package main

// 左 根 右

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//inorderTraversal 中序遍历，递归 时间O(n),空间O(n)
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    var handle func(*TreeNode)
    handle = func(root *TreeNode) {
        if root == nil {
            return
        }
        handle(root.Left)
        res = append(res, root.Val)
        handle(root.Right)
    }
    handle(root)
    return res 
}

// inorderTraversal 迭代法 时间O(n),空间O(n)
func inorderTraversal(root *TreeNode) []int {
   stack := []*TreeNode{}
   node := root
   res := []int{}

   for node != nil || len(stack) > 0 {
        // 先把左节点都入栈，左 根 右 顺序
       for node != nil {
           stack = append(stack, node)
           node = node.Left
        }
        // 出栈，将节点值加入结果
        node = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, node.Val)
        // 右边节点
        node = node.Right
   }
   return res
}


func main() {
	
}