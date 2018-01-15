package fifty

import "leetcode/algorithms"

func reverseKGroup(head *algorithms.ListNode, k int) *algorithms.ListNode {
	node := &algorithms.ListNode{}
	current := node
	index := 0
	arr := []int{}

	// 如果数组元素个数等于k个 将数组中的元素倒置放入链表
	// 否则 顺序放入链表
	var assignVale = func() {
		len := len(arr)
		if len == k {
			for i := len - 1; i >= 0; i-- {
				current.Next = &algorithms.ListNode{}
				current = current.Next
				current.Val = arr[i]
			}
		} else {
			for i := 0; i < len; i++ {
				current.Next = &algorithms.ListNode{}
				current = current.Next
				current.Val = arr[i]
			}
		}
	}

	for i := head; i != nil; i = i.Next {
		if index%k == 0 {
			assignVale()
			arr = []int{}
		}
		arr = append(arr, i.Val)

		index++
	}

	// 循环结束 剩余元素处理
	if len(arr) > 0 {
		assignVale()
	}

	return node.Next
}
