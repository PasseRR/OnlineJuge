package fifty

// 普通查询直接通过了
func normalSearch(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}

	return -1
}
