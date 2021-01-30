package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}
 

 // mergeTwoLists1 递归方法
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val > l2.Val {
        l2.Next = mergeTwoLists1(l1, l2.Next)
        return l2
    } else {
        l1.Next = mergeTwoLists1(l1.Next, l2)
        return l1
    }
}

 // mergeTwoLists2 for循环遍历
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
    start := &ListNode{}
    prev := start

    // 终止条件 l1,l2都是nil
    for l1 != nil || l2 != nil {
        if l1 == nil && l2 != nil {
            prev.Next = l2
            l2 = l2.Next
        }
        if l2 == nil && l1 != nil {
            prev.Next = l1
            l1 = l1.Next
        }
        if l1 != nil && l2 != nil {
            if l1.Val > l2.Val {
                prev.Next = l2
                l2 = l2.Next
            } else {
                prev.Next = l1
                l1 = l1.Next
            }
        }
        prev = prev.Next
    }
    return start.Next
}

func main() {
	res := mergeTwoLists1(
        &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
        &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
    )

    // res := mergeTwoLists2(
    //     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
    //     &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
    // )

	fmt.Print(res.Next.Next.Next)
}