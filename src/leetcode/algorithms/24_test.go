package algorithms

import "testing"

func TestSwapPairs(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	printNode(swapPairs(node))
}
