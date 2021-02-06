// 前K个高频元素
package main

// 导入heap包
import "container/heap"
 
func topKFrequent(nums []int, k int) []int {
	// map记录出现次数
    tmp := map[int]int{}
    for _, num := range nums {
        tmp[num]++
    }
    h := &IHeap{}
    heap.Init(h)
    // 遍历map
    for key, value := range tmp {
    	// 入堆
        heap.Push(h, [2]int{key, value})
        // 长度超过K,出堆最小元素
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    // 此时堆里都是前k个最大元素，获取结果
    ret := make([]int, k)
    for i := 0; i < k; i++ {
        ret[k - i - 1] = heap.Pop(h).([2]int)[0]
    }
    return ret
}

// 通过小顶堆实现，列表里是一个2个数字的列表嵌套
type IHeap [][2]int
// 返回长度
func (h IHeap) Len() int           { return len(h) }
// 小的值
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
// 交换
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// 入堆
func (h *IHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}
// 出堆，得到最小元素
func (h *IHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


func main() {
	
}