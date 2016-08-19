package leetcode

import (
	"math"
	"strconv"
)

func getNumber(l *ListNode) float64 {
	var result float64 = 0
	temp := l
	for i := 0; ; i++ {
		result += float64(temp.Val) * math.Pow10(i)
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := getNumber(l1) + getNumber(l2)
	sumStr := strconv.FormatFloat(sum, 'f', -1, 64)
	node := ListNode{}
	next := &node
	for i := len(sumStr) - 1; i > -1; i-- {
		val, _ := strconv.Atoi(sumStr[i : i+1])
		next.Val = val
		next.Next = &ListNode{}
		next = next.Next
	}
	return &node
}

func initNode(nums []int) *ListNode {
	node := ListNode{}
	next := &node
	for _, i := range nums {
		next.Val = i
		next.Next = &ListNode{}
		next = next.Next
	}

	return &node
}
