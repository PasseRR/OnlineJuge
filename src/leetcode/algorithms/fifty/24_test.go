package fifty

import (
	"leetcode/algorithms"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	node := &algorithms.ListNode{
		Val: 1,
		Next: &algorithms.ListNode{
			Val: 2,
			Next: &algorithms.ListNode{
				Val: 3,
			},
		},
	}

	algorithms.PrintNode(swapPairs(node))
}
