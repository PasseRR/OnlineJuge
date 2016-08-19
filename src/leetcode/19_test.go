package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	result := removeNthFromEnd(node, 3)
	printNode(result)
}
