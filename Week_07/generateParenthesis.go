// 括号生成
package main

// generateParenthesis。。。
// 递归回溯
func generateParenthesis(n int) []string {
    res := []string{}
    return dfs(n, 0, 0, "", res)
}

func dfs(n, left, right int, str string, res []string) []string {
    if len(str) == 2*n {
        res = append(res, str)
        return res
    }
    // 左括号小于n，可以添加
    if left < n {
        str += "("
        res = dfs(n, left+1, right, str, res)
        str = str[:len(str) - 1]
    }
    // 右括号小于左括号，可以添加
    if right < left {
        str += ")"
        res = dfs(n, left, right+1, str, res)
        str = str[:len(str) - 1]
    }
    return res
}

func main() {
    
}