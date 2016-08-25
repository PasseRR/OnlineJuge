package leetcode

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化一个空节点 最后返回链表的时候不要这个节点
	node := &ListNode{}
	current := node
	node1, node2 := l1, l2

	// 依次取出链表中的元素比较 直到其中一个链表为空
	for node1 != nil && node2 != nil {
		current.Next = &ListNode{}
		current = current.Next

		if node1.Val > node2.Val {
			current.Val = node2.Val
			node2 = node2.Next
		} else {
			current.Val = node1.Val
			node1 = node1.Next
		}
	}

	// 若其中一个链表还有元素
	if node1 != nil {
		current.Next = node1
	}

	if node2 != nil {
		current.Next = node2
	}

	// 去掉int默认为0的情况
	return node.Next
}
