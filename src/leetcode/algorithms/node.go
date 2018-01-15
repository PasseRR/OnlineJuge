package algorithms

import (
	"fmt"
)

// 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印单链表
func PrintNode(node *ListNode) {
	for i := node; i != nil; i = i.Next {
		if i.Next != nil {
			fmt.Printf("%v->", i.Val)
		} else {
			fmt.Printf("%v", i.Val)
		}
	}
	fmt.Printf("\n")
}

// 打印单链表数组
func PrintNodeArray(lists []*ListNode) {
	for _, item := range lists {
		PrintNode(item)
	}
}
