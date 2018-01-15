package page1

import (
	"leetcode/algorithms"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	node := &algorithms.ListNode{
		Val: 1,
		Next: &algorithms.ListNode{
			Val: 2,
			Next: &algorithms.ListNode{
				Val: 3,
			},
		},
	}

	result := removeNthFromEnd(node, 3)
	algorithms.PrintNode(result)
}
