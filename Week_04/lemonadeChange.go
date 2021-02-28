// 柠檬水找零
package main

// lemonadeChange ....
// 两个指针记录5， 10的数量，通过，5，10，20的计算，每次循环判断数量小于0则false
func lemonadeChange(bills []int) bool {
    f, t := 0, 0
    for i := 0; i < len(bills); i++ {
        if bills[i] == 5 {f += 1}
        if bills[i] == 10 {
            t += 1
            f -= 1
        }
        if bills[i] == 20 {
            if t > 0 {
                t -= 1
                f -= 1
            } else {f -= 3}
        }
        if f < 0 || t < 0 {return false}
    }
    return true
}

func main() {
    
}