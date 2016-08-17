package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	node := &ListNode{}
	current := node
	var previous, next int
	count := 0
	for i := head; i != nil; i = i.Next {
		if count%2 == 0 {
			previous = i.Val
		} else {
			next = i.Val
			current.Next = &ListNode{}
			current = current.Next
			current.Val = next
			current.Next = &ListNode{}
			current = current.Next
			current.Val = previous
		}
		count++
	}

	if count%2 == 1 {
		current.Next = &ListNode{}
		current = current.Next
		current.Val = previous
	}

	return node.Next
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	swapPairs(node)
}
