package fifty

import (
	"leetcode/algorithms"
	"testing"
)

func TestMergeKLists(t *testing.T) {
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

	lists := []*algorithms.ListNode{node1, node2}

	algorithms.PrintNode(mergeKLists(lists))
}
