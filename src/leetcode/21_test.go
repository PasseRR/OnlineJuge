package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
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

	node := mergeTwoLists(node1, node2)
	printNode(node)
}
