package fifty

import (
	"leetcode/algorithms"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	node := &algorithms.ListNode{
		Val: 1,
		Next: &algorithms.ListNode{
			Val: 2,
			Next: &algorithms.ListNode{
				Val: 3,
				Next: &algorithms.ListNode{
					Val: 4,
					Next: &algorithms.ListNode{
						Val: 5,
						Next: &algorithms.ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	algorithms.PrintNode(reverseKGroup(node, 2))
	algorithms.PrintNode(reverseKGroup(node, 4))
}
