package page1

/**
 * 中位数：一个有序数组中最中间的数，若数组元素个数为奇数，则为最中间那个元素；
 * 否则，为最中间两个元素的的平均值
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	len := len1 + len2
	if len == 0 {
		return float64(0)
	}
	nums := make([]int, len)
	i, j, k := 0, 0, 0
	for ; i < len1 && j < len2; k++ {
		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}

	for ; i < len1; k++ {
		nums[k] = nums1[i]
		i++
	}

	for ; j < len2; k++ {
		nums[k] = nums2[j]
		j++
	}
	var result float64
	if len%2 == 1 {
		result = float64(nums[len/2])
	} else {
		result = float64((nums[len/2-1] + nums[len/2])) / 2.0
	}

	return result
}
