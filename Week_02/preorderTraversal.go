// 二叉树的前序遍历
package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//preorderTraversal 前序遍历，递归
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


func main() {
	
}