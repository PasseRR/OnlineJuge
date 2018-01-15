package page1

func removeElement(nums []int, val int) int {
	count := 0
	// 双向遍历
	// 从尾部遍历值为val的
	// 从头遍历值非val的
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i] == val {
			if nums[j] != val {
				// 若头部为val 尾部非val交换两个元素位置
				nums[i], nums[j] = nums[j], nums[i]
				i++
				count++
			} else {
				// 若尾部和头部均为val 尾部继续向前遍历
				j--
			}
		} else {
			// 若头部为非val 数组长度增加 直接遍历下一个
			i++
			count++
		}
	}

	return count
}
