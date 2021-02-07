学习笔记

## 树，二叉树，二叉搜索树

### 树

树的深度为层次数

1. 树是元素的集合。

2. 该集合可以为空。这时树中没有元素，我们称树为空树 (empty tree)。

3. 如果该集合不为空，那么该集合有一个根节点，以及0个或者多个子树。根节点与它的子树的根节点用一个边(edge)相连

树的实现：
每个节点储存元素和多个指向子节点的指针。然而，子节点数目是不确定的。一个父节点可能有大量的子节点，而另一个父节点可能只有一个子节点，而树的增删节点操作会让子节点的数目发生进一步的变化。这种不确定性就可能带来大量的内存相关操作，并且容易造成内存的浪费

节点模板
``` javascript
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
```

### 二叉树，二叉搜索树

二叉树的每个节点最多只能有2个子节点

二叉搜索树可以方便的实现搜索算法。在搜索元素x的时候，我们可以将x和根节点比较:

1. 如果x等于根节点，那么找到x，停止搜索 (终止条件)

2. 如果x小于根节点，那么搜索左子树

3. 如果x大于根节点，那么搜索右子树

二叉搜索树所需要进行的操作次数最多与树的深度相等。n个节点的二叉搜索树的深度最多为n，最少为log(n)

### 二叉树遍历

时间，空间均为O（n），迭代法，递归法

前序（先序）： 根左右

中序：左根右

后序：左右根

## 堆

分为大顶堆，小顶堆

查找最大或最小：O(1)

删除最大或最小：O（logn）

插入： O（logn）或O（1）

golang 标准库，实现的是小顶堆， "container/heap"包

``` javascript
// 通过小顶堆实现，列表里是一个2个数字的列表嵌套
type IHeap []interface{}
// 返回长度
func (h IHeap) Len() int { 
	return len(h) 
}
// 小的值
func (h IHeap) Less(i, j int) bool { 
	return h[i] < h[j] 
}
// 交换
func (h IHeap) Swap(i, j int) { 
	h[i], h[j] = h[j], h[i]
}
// 入堆
func (h *IHeap) Push(x interface{}) {
    *h = append(*h, x)
}
// 出堆，得到最小元素
func (h *IHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```


### 二叉堆

1.是一颗完全二叉树，根节点都是满的，除了最后一层

2.树中任意节点的值大于等于或小于等于（二叉大顶，小顶）其子节点的值

3.二叉堆一般都是通过数组实现

4.父节点，子节点位置关系，索引为i

	左子节点：2*i+1
	右子节点：2*i+2
	父节点：floor((i-1)/2)

5.插入元素：插入堆尾部，重新维护堆，向上调整

6。删除元素：向下调整

#### golang实现二叉堆

通过数组实现一个二叉大顶堆

``` javascript
// 最大堆的定义和实现：
// 最大堆
type maxHeap struct {
	size int
	nums []int
}

// 获取父节点索引
func parent(i int) int {
	if i == 0 {
		return 0
	}
	return (i - 1) / 2
}

// 获取左节点索引
func leftChild(i int) int {
	return 2*i + 1
}

// 右节点索引
func rightChild(i int) int {
	return 2*i + 2
}

// 初始化
func NewMaxHeap() *maxHeap {
	return &maxHeap{}
}

// 获取大小
func (heap *maxHeap) GetSize() int {
	return heap.size
}

// 判断是否为空
func (heap *maxHeap) IsEmpty() bool {
	return heap.size == 0
}

// 插入元素，并向上调整
func (heap *maxHeap) SiftUp(i int) {
	// 小于则赋值
	if heap.size < len(heap.nums) {
		heap.nums[heap.size] = i
	} else { // 大于则扩容
		heap.nums = append(heap.nums, i)
	}
	// 插入的是堆尾，此时的heap.size还没更新，也就是索引
	parI := parent(heap.size)
	childI := heap.size
	// 父节点小于子节点，则交换
	for heap.nums[parI] < heap.nums[childI] {
		heap.nums[parI], heap.nums[childI] = heap.nums[childI], heap.nums[parI]
		childI = parI
		parI = parent(parI)
	}
	heap.size++
}

// 向下调整
func siftDown(heap *maxHeap, parI int) {
	var maxI int
	for {
		leftI := leftChild(parI)
		switch {
		// 左索引超出size
		case leftI+1 > heap.size:
			return

			// 左索引不超,右索引超出size,说明左索引是最后索引
		case leftI+2 > heap.size:
			if heap.nums[parI] < heap.nums[leftI] {
				heap.nums[parI], heap.nums[leftI] = heap.nums[leftI], heap.nums[parI]
			}
			return
		case heap.nums[leftI] >= heap.nums[leftI+1]:
			maxI = leftI
		case heap.nums[leftI] < heap.nums[leftI+1]:
			maxI = leftI + 1
		}

		// 比左右子节点的值都大,返回
		if heap.nums[parI] >= heap.nums[maxI] {
			return
		}

		heap.nums[parI], heap.nums[maxI] = heap.nums[maxI], heap.nums[parI]
		parI = maxI
	}
}

// 取出对中最大元素,即root节点值
func (heap *maxHeap) SiftDown() (int, error) {
	if heap.IsEmpty() {
		return 0, errors.New("maxHeap is empty.")
	}
	vTop := heap.nums[0]
	heap.size--
	heap.nums[0], heap.nums[heap.size] = heap.nums[heap.size], 0

	siftDown(heap, 0)

	return vTop, nil
}


// 查看堆中最大元素
func (heap *maxHeap) GetMax() (int, error) {
	if heap.IsEmpty() {
		return 0, errors.New("maxHeap is empty.")
	}
	return heap.nums[0], nil
}
```

## go的map分析

主要使用go语言，看了go的hashmap

golang： hashmap小结

go 的 map 底层通过哈希表实现，数组加链表的结构

### 基本结构

基本结构如下，runtime/map.go go版本 v1.11

``` javascript
// A header for a Go map.
type hmap struct {
        // Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
        // Make sure this stays in sync with the compiler's definition.
        count     int // # live cells == size of map.  Must be first (used by len() builtin)
        flags     uint8
        B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
        noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
        hash0     uint32 // hash seed

        buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
        oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
        nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

        extra *mapextra // optional fields
}

```

1.这个hash结构使用的是一个可扩展哈希的算法，由hash值mod当前hash表大小决定某一个值属于哪个桶，而hash表大小是2的指数，即上面结构体中的2^B。

2.每次扩容，会增大到上次大小的两倍。结构体中有一个buckets和一个oldbuckets是用来实现增量扩容的。正常情况下直接使用buckets，而oldbuckets为空。如果当前哈希表正在扩容中，则oldbuckets不为空，并且buckets大小是oldbuckets大小的两倍

捅结构：
``` javascript
// A bucket for a Go map.
type bmap struct {
        // tophash generally contains the top byte of the hash value
        // for each key in this bucket. If tophash[0] < minTopHash,
        // tophash[0] is a bucket evacuation state instead.
        tophash [bucketCnt]uint8
        // Followed by bucketCnt keys and then bucketCnt values.
        // NOTE: packing all the keys together and then all the values together makes the
        // code a bit more complicated than alternating key/value/key/value/... but it allows
        // us to eliminate padding which would be needed for, e.g., map[int64]int8.
        // Followed by an overflow pointer.
}
```

每个桶bucket最多存放8个kv对，多于8个，会申请一个新的bucket，并将它与之前的bucket链起来

当两个不同的 key 落在同一个桶中，也就是发生了哈希冲突。冲突的解决手段是用链表法：在 bucket 中，从前往后找到第一个空位。这样，在查找某个 key 时，先找到对应的桶，再去遍历 bucket 中的 key

### 增量扩容

扩容过程中会有两张表，oldtable 旧哈希表， newtable 新哈希表

1.根据结构设置，哈希表的大小是2^B，扩容之后的大小为2^(B+1)，每次扩容都变为原来大小的两倍

2.会建立一个新的2倍大小的hash表，将旧的buckets搬到新表

3.一般旧表hash值不等于新表hash，需要重新计算hash到新表，旧的表的buckets不会删除，加一个已删除标记

4.增量扩容，不会一次完成，每次搬移1-2个键值对，由于这个工作是逐渐完成的，这样就会导致一部分数据在old table中，一部分在new table中，所以对于hashmap的insert,remove,lookup操作的处理逻辑产生影响。只有当所有的bucket都从旧表移到新表之后，才会将oldbucket释放掉

PS： 为什么是增量， 为了缩短map容器的响应时间，元素很多的话，不增量就会卡住，响应时间过长

5.扩容因子：如果grow的太频繁，会造成空间的利用率很低，如果很久才grow，会形成很多的overflow buckets，查找的效率也会下降

作者测试：6.5适中，table中元素的个数大于table中能容纳的元素的个数

### kv的存储及查找

1、将key的类型选用hash算法得到key的hash值

2、hash值的的低位，作为buckets数组的index索引，用于找到key所在的bucket数组

3.hash的高8位存储到tophash，tophash在bucket结构里

PS：存储过程中，key放在一起。value放在一起，方便字节对齐（底层cpu访问数据效率）避免浪费存储空间

部分代码： mapaccess1方法
``` javascript
    // 不同类型 key 使用的 hash 算法在编译期确定
	alg := t.key.alg
	// 计算哈希值，并且加入 hash0 引入随机性
	hash := alg.hash(key, uintptr(h.hash0))
	// 比如 B=5，那 m 就是31，二进制是全 1
	// 求 bucket num 时，将 hash 与 m 相与，
	// 达到 bucket num 由 hash 的低 8 位决定的效果
	m := bucketMask(h.B)
	// b 就是 bucket 的地址
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
	// oldbuckets 不为 nil，说明发生了扩容
	if c := h.oldbuckets; c != nil {
		// 如果不是同 size 扩容（看后面扩容的内容）
		// 对应条件 1 的解决方案
		if !h.sameSizeGrow() {
			// 新 bucket 数量是老的 2 倍
			m >>= 1
		}
		// 求出 key 在老的 map 中的 bucket 位置
		oldb := (*bmap)(add(c, (hash&m)*uintptr(t.bucketsize)))
		// 如果 oldb 没有搬迁到新的 bucket
		// 那就在老的 bucket 中寻找
		if !evacuated(oldb) {
			b = oldb
		}
	}
	// 计算出高 8 位的 hash
	// 相当于右移 56 位，只取高8位
	top := tophash(hash)
	//开始寻找key
	for ; b != nil; b = b.overflow(t) {
		// 遍历 8 个 bucket
		for i := uintptr(0); i < bucketCnt; i++ {
			// tophash 不匹配，继续
			if b.tophash[i] != top {
				continue
			}
			// tophash 匹配，定位到 key 的位置
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			// key 是指针
			if t.indirectkey {
				// 解引用
				k = *((*unsafe.Pointer)(k))
			}
			// 如果 key 相等
			if alg.equal(key, k) {
				// 定位到 value 的位置
				v := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
				// value 解引用
				if t.indirectvalue {
					v = *((*unsafe.Pointer)(v))
				}
				return v
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0])
}
```

查找流程：

1、先找oldtable，旧的hash表，如果hash表已经被添加删除标记，则去新的找

2.先比较key的hash值高8位在tophash[i]值是否相等，顺序比较，得到i

3.找到key的hash低位对应的bucket

4.如果相等则再比较bucket的第i个的key与所给的key是否相等

5.如果没有，则在overflow里继续寻找，桶溢出链表


### map插入

1.初始化校验，根据key计算hash，根据低位找对对应的buckets数组

部分代码：
``` javascript
func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	//判断 hmap 是否已经初始化（是否为 nil）
    if h == nil {
		panic(plainError("assignment to entry in nil map"))
	}
	//...
    //判断是否并发读写 map，若是则抛出异常
	if h.flags&hashWriting != 0 {
		throw("concurrent map writes")
	}
    //根据 key 的不同类型调用不同的 hash 方法计算得出 hash 值
	alg := t.key.alg
	hash := alg.hash(key, uintptr(h.hash0))
    //设置 flags 标志位，表示有一个 goroutine 正在写入数据。因为 alg.hash 有可能出现 panic 导致异常
	h.flags |= hashWriting
    //判断 buckets 是否为 nil，若是则调用 newobject 根据当前 bucket 大小进行分配
    //初始化时没有初始 buckets，那么它在第一次赋值时就会对 buckets 分配
	if h.buckets == nil {
		h.buckets = newobject(t.bucket) // newarray(t.bucket, 1)
	}	
}
```

2.根据table中元素的个数，判断是否正在扩容。

3.如果在扩容中：
	oldbucket是被冻结的，查找时会在oldbucket中查找，但不会在oldbucket中插入数据。如果在oldbucket是找到了相应的key，做法是将它迁移到新bucket后加入evalucated标记

4.在bucket中，查找空闲的位置，如果已经存在需要插入的key，更新其对应的value。

5.如果对应的bucket已经full，重新申请新的bucket作为overbucket。

6.将key/value pair插入到bucket中

部分代码：
``` javascript
//根据低八位计算得到 bucket 的内存地址
	bucket := hash & bucketMask(h.B)
	//判断是否正在扩容，若正在扩容中则先迁移再接着处理
	if h.growing() {
		growWork(t, h, bucket)
	}
	//计算并得到 bucket 的 bmap 指针地址
	b := (*bmap)(unsafe.Pointer(uintptr(h.buckets) + bucket*uintptr(t.bucketsize)))
	//计算 key hash 高八位用于查找 Key
	top := tophash(hash)
	var inserti *uint8
	var insertk unsafe.Pointer
	var val unsafe.Pointer
	for {
		//迭代 buckets 中的每一个 bucket（共 8 个）
		for i := uintptr(0); i < bucketCnt; i++ {
			//对比 bucket.tophash 与 top（高八位）是否一致
			if b.tophash[i] != top {
				//若不一致，判断是否为空槽
				if b.tophash[i] == empty && inserti == nil {
					//有两种情况，第一种是没有插入过。第二种是插入后被删除
					inserti = &b.tophash[i]
					insertk = add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
					//把该位置标识为可插入 tophash 位置,这里就是第一个可以插入数据的地方
					val = add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
				}
				continue
			}
			//若是匹配（也就是原本已经存在），则进行更新。最后跳出并返回 value 的内存地址
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey {
				k = *((*unsafe.Pointer)(k))
			}
			if !alg.equal(key, k) {
				continue
			}
			// already have a mapping for key. Update it.
			if t.needkeyupdate {
				typedmemmove(t.key, k, key)
			}
			val = add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.valuesize))
			goto done
		}
		//判断是否迭代完毕，若是则结束迭代 buckets 并更新当前桶位置
		ovf := b.overflow(t)
		if ovf == nil {
			break
		}
		b = ovf
	}
	//若满足三个条件：触发最大 LoadFactor 、存在过多溢出桶 overflow buckets、没有正在进行扩容。就会进行扩容动作（以确保后续的动作）
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}

```

PS：
插入是数据追加方式：
1.先找到对应的key，则更新对应的value并结束遍历

2.先找到对应的空位，则直接在这个位置插入，后面的都会覆盖并结束遍历，因为是顺序遍历，buckets数组后面的会被覆盖掉

3.然后就是只要在某个bucket中找到第一个空位，就会将key/value插入到这个位置。也就是位置位于bucket前面的会覆盖后面的(类似于存储系统设计中做删除时的常用的技巧之一，直接用新数据追加方式写，新版本数据覆盖老版本数据)

不过这也意味着做删除时必须完全的遍历bucket所有溢出链，将所有的相同key数据都删除。

所以目前map的设计是为插入而优化的，删除效率会比插入低一些

### map删除

删除某个key的操作与分配类似，由于hashmap的存储结构是数组+链表，所以真正删除key仅仅是将对应的slot设置为empty，并没有减少内存

delete(intMap, 1)

所以过程给查找类似，把tophash值设成empty，并不是真正释放内存