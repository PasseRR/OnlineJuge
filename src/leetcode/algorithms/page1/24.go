package page1

import "leetcode/algorithms"

func swapPairs(head *algorithms.ListNode) *algorithms.ListNode {
	node := &algorithms.ListNode{}
	current := node
	var previous, next int
	count := 0
	for i := head; i != nil; i = i.Next {
		if count%2 == 0 {
			previous = i.Val
		} else {
			next = i.Val
			current.Next = &algorithms.ListNode{}
			current = current.Next
			current.Val = next
			current.Next = &algorithms.ListNode{}
			current = current.Next
			current.Val = previous
		}
		count++
	}

	if count%2 == 1 {
		current.Next = &algorithms.ListNode{}
		current = current.Next
		current.Val = previous
	}

	return node.Next
}
