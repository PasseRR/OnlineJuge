package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 单链表的长度
	len := 0
	for node := head; node != nil; node = node.Next {
		len++
	}
	// 特殊处理移除头节点
	if len == n {
		return head.Next
	} else {
		for i, node := 0, head; node.Next != nil; node = node.Next {
			if len-i-1 == n {
				node.Next = node.Next.Next
				break
			}
			i++
		}
		return head
	}
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

	result := removeNthFromEnd(node, 3)
	for i := result; i != nil; i = i.Next {
		fmt.Printf("%v->", i.Val)
	}
}
