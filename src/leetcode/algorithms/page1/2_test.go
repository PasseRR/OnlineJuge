package page1

import (
	"leetcode/algorithms"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l2 := &algorithms.ListNode{
		Val: 1,
	}

	nums := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	algorithms.PrintNode(addTwoNumbers(initNode(nums), l2))
}

func initNode(nums []int) *algorithms.ListNode {
	node := algorithms.ListNode{}
	next := &node
	for _, i := range nums {
		next.Val = i
		next.Next = &algorithms.ListNode{}
		next = next.Next
	}

	return &node
}
