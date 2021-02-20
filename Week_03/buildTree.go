// 从前序中序构造二叉树
package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// buildTree 递归法
// 1.前序遍历的第一个节点，就是根节点，根左右，根节点后面跟的是左子树，然后是右子树
// 2.中序遍历的根节点在中间，根节点的位置左边就是左子树，右边就是右子树
// 3.因为都是一个树，所以中序和前序的左子树长度相等，同理右子树，所以中序的根节点以前长度在前序根节点以后的长度就是左子树
// 4.重复子问题就是，前序得到根节点，中序找到根节点，根节点之前就是左子树，左子树出左节点，根节点之后就是右子树，右子树出右节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 终止条件，前序长度为0，说明已经递归完
    if len(preorder)==0{
        return nil
    }
    // 前序的第一个节点就是根节点
    root:=&TreeNode{
        Val:preorder[0],
    }
    // 找到根节点在中序的位置
    var num int
    for k, v := range inorder{
        if v == preorder[0]{
            num = k
            break
        }
    }
    // 处理左右子树，前序是出去第一个元素，包含num的元素，中序是包含第一个元素，不包含num的元素
    root.Left = buildTree(preorder[1:num+1],inorder[:num])
    root.Right = buildTree(preorder[num+1:],inorder[num+1:])
    return root
}

func main() {
    
}