学习笔记

### lru cache

Lru Cache算法就是LeastRecentlyUsed，意思就是最近最少使用，当缓存满了的时候，不经常使用的就直接删除，挪出空间来缓存新的对象

LRU 缓存机制可以通过哈希表辅以双向链表实现，我们用一个哈希表和一个双向链表维护所有在缓存中的键值对。

双向链表按照被使用的顺序存储了这些键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。

哈希表即为普通的哈希映射（HashMap），通过缓存数据的键映射到其在双向链表中的位置

首先使用哈希表进行定位，找出缓存项在双向链表中的位置，随后将其移动到双向链表的头部，即可在 O(1) 的时间内完成 get 或者 put 操作

本质是一个map集合，然后把map集合里的元素，放到了一个双向链表里，常用的放前面，不常用放后面

golang实现LRU cache

``` javascript
// 结构体
type LRUCache struct {
    size int        // 大小
    capacity int    // 容量
    cache map[int]*DLinkedNode  //  hashmap，双向链表，根据key值查找node
    head, tail *DLinkedNode     // 头指针， 尾指针
}
// 双向链表
type DLinkedNode struct {
    key, value int // key值，value值
    prev, next *DLinkedNode // 先前指针，后序指针
}
// 初始化链表
func initDLinkedNode(key, value int) *DLinkedNode {
    return &DLinkedNode{
        key: key,
        value: value,
    }
}
// 初始化一个lru cache
func Constructor(capacity int) LRUCache {
    // 初始化结构体
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0, 0),
        tail: initDLinkedNode(0, 0),
        capacity: capacity,
    }
    // 头指针指向尾指针
    l.head.next = l.tail
    // 尾指针指向头指针
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    // 根据key取出值
    node := this.cache[key]
    // 放入头部
    this.moveToHead(node)
    // 返回值
    return node.value
}

// 放入元素
func (this *LRUCache) Put(key int, value int)  {
    // 如果key不是在cache里
    if _, ok := this.cache[key]; !ok {
        // 创建一个新的节点
        node := initDLinkedNode(key, value)
        // 将新的节点放到head
        this.cache[key] = node
        this.addToHead(node)
        // 大小加1
        this.size++
        // 如果数量大于容量
        if this.size > this.capacity {
            // 删除尾指针的node'
            removed := this.removeTail()
            // 删除字典中的key数据
            delete(this.cache, removed.key)
            this.size--
        }
    } else {
        // 如果在字典里，则再次放入的话提高了了频次，移动到头部
        node := this.cache[key]
        node.value = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    // 头部添加，node的前面指向head
    node.prev = this.head
    // node的next指向原先head的next
    node.next = this.head.next
    // head的next的前面指向node
    this.head.next.prev = node
    // head的next指向node
    this.head.next = node
}
// 删除node节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
    // node前面的节点的next指向node的next
    node.prev.next = node.next
    // node的next的前面指向node的前面
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    // 先删除node
    this.removeNode(node)
    // 再把node加入到前面
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
    // 这个时候新的node已经加到前面了，所以要删除的是tail前面的prev
    node := this.tail.prev
    // 删除
    this.removeNode(node)
    return node
}
```

### 布隆过滤器

布隆过滤器是二进制向量数据结构，它具有很好的空间和时间效率，被用来检测一个元素是不是集合中的一个成员

判定可能已存在和绝对不存在两种情况。

如果检测结果为是，该元素不一定在集合中

但如果检测结果为否，该元素一定不在集合中


应用场景：

ip黑名单，爬虫url过滤， 垃圾邮件过滤，去重


实现步骤：

1、开辟一个长度为m的bit数组，可以用文件来实现

2、通过hash函数计算得到位置值，并将数组里对应的位置的二进制置为1

3、查找时，通养通过hash计算得到位置，该位置的如果都是1，可能存在，如果有0，则一定不存在

优点：

有很好的空间和时间效率

不需要存储元素本身，在某些对保密要求非常严格的场合有优势

存储空间和插入查询都是常数复杂度

缺点：

误判率会随元素的增加而增加

不能从布隆过滤器中删除元素


redis 4.0以上支持布隆过滤器

### 排序

### 快速排序

分成两个数组，对每个数组进行递归排序，不断递归，最后得到排好序的 数组

先处理排序，再递归

``` javascript
func quickSort(q []int, l, r int) {
    // 递归终止条件
    if l >= r {
        return
    }
    x := q[(l+r)>>1] // 确定分界点，从中间开始
    i, j := l-1, r+1 // 两个指针，因为do while要先自增/自减
    // ，大的放右边，小的放左边
    for i < j {
        for {
            i++
            // 碰到大于的就停止
            if q[i] >= x { break }
        }
        for {
            j--
            // 碰到小于的就停止
            if q[j] <= x { break}
        }
        if i < j { // swap 两个元素
            q[i], q[j] = q[j], q[i]
        }
    }
    // 递归处理左右两段
    quickSort(q, l, j) 
    quickSort(q, j+1, r)
}

```

时间：O(nlogn)
空间：O(n) 
稳定

### 归并排序

分治思想与快排相反，先排，再合并，先递归在合并

``` javascript

func merge_sort(nums []int, left, right int) {
    if left >= right { return }
    // 找中间位置
    mid := (left + right) >> 1
    // 左右排序
    merge_sort1(nums, left, mid)
    merge_sort1(nums, mid+1, right)
    //合并
    tmp := []int{}
    // 左右的起始位置
    i, j := left, mid+1
    for i <= mid && j <= right {
        // 比较两个指针的位置，放入新数组
        if nums[i] <= nums[j] {
            tmp = append(tmp, nums[i])
            i++
        } else {
            tmp = append(tmp, nums[j])
            j++
        }
    }
    // 没排完的，放到后面
    if i <= mid {
        tmp = append(tmp, nums[i:mid+1]...)
    } else { //如果使用...运算符，可以将一个切片的所有元素追加到另一个切片里
        tmp = append(tmp, nums[j:right+1]...)
    }
    // 放到nums里
    copy(nums[left:right+1], tmp)
}
```
时间： O(nlogn)
空间：O(n) 临时数组加递归深度
稳定

### 堆排序

将数组生成大顶堆，或者小顶堆，再逐步取出堆顶元素

``` javascript
func heap_sort(nums []int) []int {
    lens := len(nums) - 1
    // 建堆 O(n) lens/2后面都是叶子节点，不需要向下调整
    for i := lens/2; i >= 0; i -- { 
        down(nums, i, lens)
    }
    for j := lens; j >= 1; j -- {
        // 将最大值放到后面
        nums[0], nums[j] = nums[j], nums[0]
        // 最后一位已经是最大值，调整之前的值
        lens --
        down(nums, 0, lens)
    }
    return nums
}
//O(logn)大根堆，如果堆顶节点小于叶子，向下调整
func down(nums []int, i, lens int) { 
    // 当前值是最大值
    max := i 
    // 如果他的左节点大于最大值
    if i<<1+1 <= lens && nums[i<<1+1] > nums[max] {
        // 则最大值等于子节点
        max = i<<1+1
    }  
    // 如果他的右节点大于最大值
    if i<<1+2 <= lens && nums[i<<1+2] > nums[max] {
        max = i<<1 + 2
    }
    // 如果最大值发生变化，则交换两个位置
    if max != i {
        nums[max], nums[i] = nums[i], nums[max]
        // 递归向下调整
        down(nums, max, lens)
    }
}
```

时间: O(nlogn)
空间：O(1）
不稳定

### 选择排序

当前循环的值为最小值，内层循环遍历后面的，后面的小于最小值，找到一个最小的，最后交换位置

``` javascript
func select_sort(q []int) {
    for i := 0; i < len(q) - 1; i++ {
        minIndex := i
        for j := i;j < len(q); j++ {
            if q[j] < q[minIndex] {
                minIndex = j
            }
        }
        q[i], q[minIndex] = q[minIndex], q[i]
    }
}
```
时间复杂度：O(n ^2)
空间复杂度:O(1)
不稳定

### 冒泡排序

两两交换，保证最后一个是最大值，然后第二次遍历就到最后一个的前面截止

``` javascript
func bubbleSort(q []int) {
    n := len(q)
    for i := 0; i < n - 1; i++ {
        exchange := false
        for j := 0; j < n-1-i; j++ {
            if q[j] > q[j+1] {
                q[j], q[j+1] = q[j+1], q[j]
                exchange = true
            }
        }
        if !exchange {
            break
        }
    }

}
```

时间复杂度：O(n ^2)
空间复杂度:O(1)
稳定

### 插入排序

``` javascript
func insertSort(nums []int) []int {
    n := len(nums)
    for i := 1; i < n; i++ {
        tmp := nums[i]
        j := i - 1
        //左边比右边大
        for j >= 0 && nums[j] > tmp { 
            //右边的数就等于前一个数，最后大于tmp的都被j+1赋值，相当于右移
            nums[j+1] = nums[j] 
            j--                 //到前一个数
        }
        // 上一步比tmp大的 都右移过去之后，小于tmp的右边
        nums[j+1] = tmp
    }
    return nums
}
```

时间复杂度：O(n ^2)
空间复杂度:O(1)
稳定

### 希尔排序

插入排序的优化，长度/2，再次/4

``` javascript
func shell_sort(q []int) {
    // 先排n/2，后面的排好，再排n/4
    for k := len(q) / 2; k > 0; k /= 2 {
        // 从k开始往后，插入排序
        for i := k; i < len(q); i++ {
            tmp := q[i]
            j := i - k
            for j >= 0 && tmp < q[j] {
                q[j+k] = q[j]
                j -= k
            }
           q[j+k] = tmp
        }
    }
}
```

时间复杂度: O(nlogn)
空间复杂度: O(1)
不稳定

### 计数排序

找出最大值即可，然后从1到最大值生成map，有值记为1，没有的记为0

然后继续0到最大值，遍历，从字典取值，就是按顺序取的，按顺序放到数组里

``` javascript
func qSort(q []int) {
    v := [10001]int{}
    for i := 0; i < len(q); i++{
        v[5000+q[i]]++
    }
    idx := 0
    for i := 0; i < 10001; i++ {
        for v[i] > 0 {
            q[idx] = i - 5000
            idx++
            v[i]--
        }
    }
}
```

时间复杂度:O(n+k)
空间复杂度:O(k)
稳定排序

### 桶排序

``` javascript
func bin_sort(li []int, bin_num int) {
    // 找最大值，找最小值
    min_num, max_num := li[0], li[0]
    for i := 0; i < len(li); i++ {
        if min_num > li[i] {
            min_num = li[i]
        }
        if max_num < li[i] {
            max_num = li[i]
        }
    }
    // 捅数，生成捅
    bin := make([][]int, bin_num)
    for j := 0; j < len(li); j++ {
        // 计算捅的位置，减去最小值，除以 （最大值-最小值+1）/桶数
        n := (li[j] - min_num) / ((max_num - min_num + 1) / bin_num)
        // 加入捅
        bin[n] = append(bin[n], li[j])
        // ，每个捅排序，插入排序
        k := len(bin[n]) - 2
        for k >= 0 && li[j] < bin[n][k] {
            bin[n][k+1] = bin[n][k]
            k--
        }
        bin[n][k+1] = li[j]
    }
    // 排好序后给到结果
    o := 0
    for p, q := range bin {
        for t := 0; t < len(q); t++ {
            li[o] = bin[p][t]
            o++
        }
    }
}
```

时间复杂度: O(n+k)
空间复杂度: O(n+k)
稳定

### 基数排序

个位，十位，百位排序

``` javascript
func radix_sort(li []int) {
    max_num := li[0]
    for i := 0; i < len(li); i++ {
        if max_num < li[i] {
            max_num = li[i]
        }
    }
    for j := 0; j < len(strconv.Itoa(max_num)); j++ {
        bin := make([][]int, 10)
        for k := 0; k < len(li); k++ {
            n := li[k] / int(math.Pow(10, float64(j))) % 10
            bin[n] = append(bin[n], li[k])
        }
        m := 0
        for p := 0; p < len(bin); p++ {
            for q := 0; q < len(bin[p]); q++ {
                li[m] = bin[p][q]
                m++
            }
        }
    }
}

```

时间复杂度: O(kn)
空间复杂度: O(n+k)
稳定