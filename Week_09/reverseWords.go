// 翻转字符串中的单词
package main

// split，join
func reverseWords(s string) string {
     list := strings.Split(s ," ")
     var res []string
     for i:=len(list)-1;i>=0;i--{
        if len(list[i])>0{
            res = append(res,list[i])
        }
     }
     s = strings.Join(res," ")
     return s
}

// 不使用split，join
func reverseWords(s string) string {
    tmp := []string{}
    single := ""
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            if single != "" {
                tmp = append(tmp, single)
            }
            single = ""
        } else {
            single += string(s[i])
            if i == len(s) - 1 {
                tmp = append(tmp, single)
            }
        }
    }
    reverseAll(tmp)
    res := ""
    for _, v := range tmp {
        res += (" " + string(v))
    } 
    return res[1:]
}

func reverseAll(s []string) {
    for i := 0; i< len(s)>>1; i++ {
        s[i], s[len(s) - i -1] = s[len(s) - i -1], s[i]
    }
}

func main() {
    
}