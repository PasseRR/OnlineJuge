package fifty

import (
	"leetcode/algorithms"
	"math"
	"strconv"
)

func getNumber(l *algorithms.ListNode) float64 {
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

func addTwoNumbers(l1 *algorithms.ListNode, l2 *algorithms.ListNode) *algorithms.ListNode {
	sum := getNumber(l1) + getNumber(l2)
	sumStr := strconv.FormatFloat(sum, 'f', -1, 64)
	node := algorithms.ListNode{}
	next := &node
	for i := len(sumStr) - 1; i > -1; i-- {
		val, _ := strconv.Atoi(sumStr[i : i+1])
		next.Val = val
		next.Next = &algorithms.ListNode{}
		next = next.Next
	}
	return &node
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
