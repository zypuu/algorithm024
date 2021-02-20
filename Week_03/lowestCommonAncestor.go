// 二叉树的最近公共祖先
package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// lowestCommonAncestor 递归方法
// 转化为最小重复子问题，根左右节点的几种情况如下：
// 1.若root中只包含p则返回p
// 2.若root中只包含q则返回q
// 3.若root中不包含p也不包含q则返回NULL
// 4.若root中同时包含p和q，则返回root（此时root为最近公共祖先）
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 不同层。，如果该节点找到了p，q则返回该节点
    if root == nil || p == root || q == root {
        return root
    }
    // 递归就是判断当前节点是不是等于p或q，不是就一直往下找，找到就返回，或者找到nil也返回
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    // 如果存在nil则向上返回另一个节点
    if left == nil {
        return right
    }
    if right == nil {
        return left
    }
    // 左右不等于nil，说明左右等于p，q，则root根就是最近公共祖先
    return root
}

// lowestCommonAncestor 迭代法
// 1.通过递归遍历，记录所有父节点
// 2.通过map记录p节点的访问路径
// 3.然后看另一节点是否有相同的父节点，有为最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 用于记录所有的父节点路径
    parent := map[int]*TreeNode{}
    // 用于记录p节点的访问路径，并判断q的相同父节点
    visited := map[int]bool{}
    // 递归遍历树，只记录父节点
    var dfs func(*TreeNode)
    dfs = func(r *TreeNode) {
        if r == nil {
            return
        }
        // 只看左右节点是否为空，因为只记录父节点
        if r.Left != nil {
            parent[r.Left.Val] = r
            dfs(r.Left)
        }
        if r.Right != nil {
            parent[r.Right.Val] = r
            dfs(r.Right)
        }
    }
    dfs(root)

    // 记录p节点的访问路径，从下向上
    for p != nil {
        visited[p.Val] = true
        p = parent[p.Val]
    }
    // 看另一节点是否有相同值，有则返回就是结果
    for q != nil {
        if visited[q.Val] {
            return q
        }
        q = parent[q.Val]
    }

    return nil
}


func main() {
	
}