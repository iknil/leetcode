package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func getListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p, head *ListNode
	u := 0
	t := 0
	for ; l1 != nil || l2 != nil; {
		t = u
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}

		_t := getListNode(t % 10, nil)
		u = t / 10
		if p == nil {
			head = _t
			p = _t
		} else {
			p.Next = _t
			p = p.Next
		}
	}

	if u > 0 && p != nil {
		p.Next = getListNode(u, nil)
	}

	return head
}

func main() {
	l1 := getListNode(2, getListNode(4, getListNode(3, nil)))
	l2 := getListNode(5, getListNode(6, getListNode(4, nil)))

	r := addTwoNumbers(l1, l2)
	for ; r != nil ; r = r.Next {
		fmt.Println(r.Val)
	}
}