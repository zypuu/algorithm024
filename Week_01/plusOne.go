package main

import "fmt"

// plusOne ...
func plusOne(digits []int) []int {
    isAdd := true
    for i := len(digits)-1;i>=0;i-- {
        if isAdd {
            digits[i] += 1
            // 超过10，则需要进1，则当前位设为0，下一位需要加一，保持isAdd
            if digits[i] >= 10 {
                digits[i] = 0
            } else {
                isAdd = false
            }
        }
    }
    if isAdd {
        // 添加末位，copy，不需要再开辟新的slice，效率比较高
        digits = append(digits, 0)
        copy(digits[1:], digits[0:])
        digits[0] = 1
    }
    return digits
}

func main() {
	res := plusOne([]int{9, 9, 9})
	fmt.Print(res)
}