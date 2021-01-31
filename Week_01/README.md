学习笔记

## 作业题
twoSum.go: 两数之和
mergeTwoLists.go: 合并两个有序链表
moveZeros.go: 移动零
plusOne.go: 加一
removeDuplicates.go: 删除数组重复项
merge.go: 合并两个有序数组
trap.go: 接雨水
rotate.go: 数组反转
myCircularDeque.go: 循环双端队列

## 总结

### 数组，链表，跳表

数组： 
prepend: O(1), append: O(1), insert：O(n), del: O(n), lookup:O(1)
链表：
prepend: O(1), append: O(1), insert：O(1), del: O(1), lookup:O(n)
跳表：
对标平衡树，和二分查找，一定要是有序序列，通过升维，添加多级索引，空间换时间来减少链表的查询时间，O(logn)
空间复杂度位O(n)

### 栈，队列

栈(stack)：先入后出
insert：O(1), del: O(1)，lookup: O(n)
队列(queue)：先入先出
insert：O(1), del: O(1), lookup: O(n)
双端队列(deque):
insert：O(1), del: O(1), lookup: O(n)