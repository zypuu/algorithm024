// 合并两个有序数组
package main

import "fmt"
import "sort"

 // merge1 for
func merge1(nums1 []int, m int, nums2 []int, n int) {
    // 
    for k:=m+n; m>0 && n>0; k-- {
        if nums2[n-1] >= nums1[m-1] {
            nums1[k-1] = nums2[n-1]
            n--
        } else {
            nums1[k-1] = nums1[m-1]
            m--
        }
    }
    // 处理第二个数组剩下的元素
    for ;n-1>=0;n-- {
        nums1[n-1] = nums2[n-1]
    }
}

 // merge2 sort
func merge2(nums1 []int, m int, nums2 []int, n int) {
    // copy减少额外开销
    copy(nums1[m:], nums2[0:])
    sort.Ints(nums1)
}


func main() {
    nums1 := []int{1, 3, 4, 0, 0, 0}
	merge1(nums1, 3, []int{2, 5, 6}, 3)

    // merge2([]int{1, 3, 4, 0, 0, 0}, 3, []int{2, 5, 6}, 3)

	fmt.Print(nums1)
}