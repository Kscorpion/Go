package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		n1 = headA
		n2 = headB
	)
	for n1 != n2 {
		if n1 == nil {
			n1 = headB
		} else {
			n1 = n1.Next
		}
		if n2 == nil {
			n2 = headA
		} else {
			n2 = n2.Next
		}
	}
	return n1
}
