package main

import "fmt"

// rotate1 整个slice copy
func rotate1(nums []int, k int)  {
    n := len(nums)
    if k > n {
        k = k%n
    }
    if k == n {
        return
    }
    // 需要额外创建新的slice
    var a []int
    a = append(a, nums[n-k:]...)
    copy(nums[k:], nums[:n-k])
    copy(nums[:k], a)
}

// rotate2 for循环单元素copy
func rotate2(nums []int, k int)  {
    n := len(nums)
    if k == n {
        return
    }
    if k > n {
        k = k%n
    }
    for i:=1;i<=k;i++{
        a := nums[n-1]
        copy(nums[1:], nums[0:])
        nums[0] = a
    }
}

// rotate3 反转数组
func rotate3(nums []int, k int)  {
    n := len(nums)
    if k > n {
        k = k%n
    }
    if k == n {
        return
    }
    //  反转整个数组
    reverseSlice(nums)
    // k之前反转
    reverseSlice(nums[:k])
    // k之后反转
    reverseSlice(nums[k:])
}

// reverseSlice ...
func reverseSlice (arr []int) {
    for i := 0; i < len(arr)/2; i++ {
        arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i] 
    }
}

func main() {
    nums := []int{1,2,3,4,5,6,7}
	rotate1(nums, 3)
	fmt.Print(nums)
}