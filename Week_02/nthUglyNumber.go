// 丑数
package main
 
func min(a,b int) (int) {
    if a>b {
        return b
    }
    return a
}

// nthUglyNumber 三指针解法
// 当一个丑数已经被* 2 * 3 * 5后，对于生成丑数已经没有用了，我们把对应指针前移一位，让下一个丑数等待被*来生成新丑数。
// 每次计算出来个三个丑数的最小的一个作为新丑数加入，然后判断是通过 * 多少得到的来后移对应指针。
// 通过123我们可以发现，我们每次生成的都是最小的丑数（保证是升序），并且每个丑数都尝试过被 *2 *3 *5（保证无遗漏）
func nthUglyNumber(n int) int {
    ugly := []int{1}
    // 初始化三个标识，2,3,5，与临时丑数
    i,j,k,tmp := 0, 0, 0, 1
    //遍历
    for idx := 1; n <= idx; idx++ {
        // 根据三个指针，计算新的三个丑数，找到最小的
        tmp = min(ugly[i]*2, min(ugly[j]*3, ugly[k]*5))
        // 是哪个*起来的指针就加一，后移      
        if (tmp == ugly[i]*2) {
            i++
        }
        if (tmp == ugly[j]*3) {
            j++
        }
        if (tmp == ugly[k]*5) {
            k++
        }
        // 将新生成的最小丑数加入数组
        ugly = append(ugly, tmp)
    }
    return ugly[n-1]
}


func main() {
	
}