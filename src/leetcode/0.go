package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(node *ListNode) {
	for i := node; i != nil; i = i.Next {
		fmt.Printf("%v->", i.Val)
	}
	fmt.Printf("\n")
}

func printNodeArray(lists []*ListNode) {
	for _, item := range lists {
		printNode(item)
	}
}
