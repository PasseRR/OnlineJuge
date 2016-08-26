package algorithms

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5}
	t.Logf("%v", findMedianSortedArrays(nums1, nums2))
}
