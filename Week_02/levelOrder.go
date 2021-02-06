// N叉树的层序遍历
package main


type Node struct {
    Val int
    Children []*Node
}
 

// levelOrder 迭代法
func levelOrder(root *Node) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    q := []*Node{root}

    for len(q) > 0 {
        l := len(q)
        levelRes := []int{}
        for l > 0 {
            node := q[0]
            q = q[1:]

            if len(node.Children) != 0 {
                for _, v := range node.Children {
                    q = append(q, v)
                }
            }
            levelRes = append(levelRes, node.Val)
            l--
        }
    res = append(res, levelRes)
    }
    return res 
}


func main() {
	
}