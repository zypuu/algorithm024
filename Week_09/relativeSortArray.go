// 数组相对排序
package main

// 类似计数排序，数字大也可以
func relativeSortArray(arr1 []int, arr2 []int) []int {
    tmp := make(map[int]int, len(arr1))
    // 计数
    for i := 0; i < len(arr1); i++ {
        tmp[arr1[i]]++
    }
    res := make([]int, 0, len(arr1))
    // 按照数组2的 顺序，将数组1的数字减减，放入结果
    for i := 0; i < len(arr2); i++ {
        for tmp[arr2[i]] > 0 {
            res = append(res, arr2[i])
            tmp[arr2[i]]--
        }
    }
    // 上面处理完的数字map里都是0，遍历剩余的数字放入一个新的数组
    last := []int{}
    for k, v := range tmp {
        for v > 0 {
            last = append(last, k)
            v--
        }
        
    }
    // 升序排序
    sort.Ints(last)
    // 放到后面
    res =append(res, last...)
    return res
}

func main() {
    
}