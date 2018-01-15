package page1

func maxArea(height []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	var min = func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}
	area := 0
	for start, end := 0, len(height)-1; start < end; {
		// 体积是高度较小的线与x轴形成的容器
		area = max(area, (end-start)*min(height[start], height[end]))
		// 寻找更大高度的线来形成更大的容器 高度会减少
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	return area
}
