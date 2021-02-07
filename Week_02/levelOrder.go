// N叉树的层序遍历
package main


type Node struct {
    Val int
    Children []*Node
}
 

// levelOrder 迭代法， 队列实现
func levelOrder(root *Node) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    // 初始化队列，将根节点加入
    q := []*Node{root}

    for len(q) > 0 {
    	// 记录当前队列的长度
        l := len(q)
        levelRes := []int{}
        // 不断出队，遍历
        for l > 0 {
        	// 队首出队
            node := q[0]
            q = q[1:]
            // 遍历节点的子节点，下一层
            if len(node.Children) != 0 {
            	// 子节点入队，队尾入队
                for _, v := range node.Children {
                    q = append(q, v)
                }
            }
            // 出队，记录值，放入结果，这是一层的结果
            levelRes = append(levelRes, node.Val)
            l--
        }
    // 将每一层的结果放入最终结果
    res = append(res, levelRes)
    }
    return res 
}


func main() {
	
}