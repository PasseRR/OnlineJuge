package page1

func jump(nums []int) int {
	// steps 总共需要跳的步数
	// jumpMaxRange 跳一步的最大范围
	// currentRange 当前steps步数能跳最大距离
	steps, jumpMaxRange, currentRange := 0, 0, 0
	for i, value := range nums {
		// i为当前需要跳的距离
		// 若当前能跳的最大距离跳不到当前i位置 表示需要多跳一步
		if i > currentRange {
			steps++
			currentRange = jumpMaxRange
		}

		// value当前可以跳的距离 i为steps步已经跳的距离
		stepRange := value + i
		// 取得跳steps步可以跳的最大范围
		if stepRange > jumpMaxRange {
			jumpMaxRange = stepRange
		}
	}

	return steps
}
