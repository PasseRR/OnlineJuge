package algorithms

func searchRange(nums []int, target int) []int {
	position := []int{-1, -1}
	for index, value := range nums {
		if value == target {
			if position[0] == -1 {
				position[0] = index
				position[1] = index
			} else {
				position[1] = index
			}
		}
	}

	return position
}
