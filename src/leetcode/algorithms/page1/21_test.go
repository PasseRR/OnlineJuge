package page1

import (
	"leetcode/algorithms"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	node1 := &algorithms.ListNode{
		Val: 1,
		Next: &algorithms.ListNode{
			Val: 3,
		},
	}

	node2 := &algorithms.ListNode{
		Val: 2,
		Next: &algorithms.ListNode{
			Val: 4,
			Next: &algorithms.ListNode{
				Val: 55,
			},
		},
	}

	node := mergeTwoLists(node1, node2)
	algorithms.PrintNode(node)
}
