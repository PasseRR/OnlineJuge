package leetcode

// 要求不申请额外的空间
// 去掉数组的重复元素 并返回去重后数组的长度
// 保证返回长度数组的元素不重复 多余的不用管
func removeDuplicates(nums []int) int {
	len := len(nums)
	if len < 2 {
		return len
	}
	count := 0
	for i := 1; i < len; i++ {
		if nums[i] != nums[i-1] {
			count++
			nums[count] = nums[i]
		}
	}

	return count + 1
}
