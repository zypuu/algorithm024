// 设计循环双端队列
package main

import "fmt"

// MyCircularDeque 队列基本结构
type MyCircularDeque struct {
    Val     []int
    K       int
}


// Constructor 新建一个队列，初始化长度为k，并记录队列的长度
func Constructor(k int) MyCircularDeque {
    a := make([]int, 0, k)
    return MyCircularDeque{Val: a, K: k}
}


// InsertFront 从前插入，判断队列是不是满的
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.Val = append(this.Val, 0)
    copy(this.Val[1:], this.Val[0:])
    this.Val[0] = value
    return true
}


// InsertLast 结尾插入，判断是不是满的
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.Val = append(this.Val, value)
    return true
}   


// DeleteFront 从前面删除，判断队列是不是空的
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.Val = this.Val[1:]
    return true
}


// DeleteLast 从后面删除 判断是否为空
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.Val = this.Val[:len(this.Val)-1]
    return true
}


// GetFront ...
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Val[0]
}


// GetRear ...
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Val[len(this.Val)-1]
}


// IsEmpty ...
func (this *MyCircularDeque) IsEmpty() bool {
    if len(this.Val) == 0 {
        return true
    }
    return false
}


// IsFull ...
func (this *MyCircularDeque) IsFull() bool {
    if len(this.Val) == this.K {
        return true
    }
    return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
}