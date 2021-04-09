// 反转字符串里的单词
package main

// 倒序添加，根据空格分隔
func reverseWords(s string) string {
     list :=  strings.Split(s ," ")
     var res []string
     for i:=len(list)-1;i>=0;i--{
        if len(list[i])>0{
            res = append(res,list[i])
        }
     }
     s =strings.Join(res," ")
     return s
}

func main() {
    
}