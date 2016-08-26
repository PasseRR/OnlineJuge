package algorithms

import "testing"

func TestMergeKLists(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}

	node2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 55,
			},
		},
	}

	lists := []*ListNode{node1, node2}

	printNode(mergeKLists(lists))
}
