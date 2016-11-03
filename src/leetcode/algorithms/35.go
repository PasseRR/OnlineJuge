package algorithms

// 二分查找
func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}

	start, end := 0, length-1
	var mid int
	for start <= end {
		mid = (end-start)/2 + start
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}

	return start
}
