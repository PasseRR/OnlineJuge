package page1

import "leetcode/algorithms"

func removeNthFromEnd(head *algorithms.ListNode, n int) *algorithms.ListNode {
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
