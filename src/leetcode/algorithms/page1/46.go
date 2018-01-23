package page1

func permute(nums []int) [][]int {
	// 排列组合结果集
	permutations := [][]int{}
	// 索引节点是否访问记录
	visited := map[int]bool{}
	dfsPermute(nums, []int{}, visited, &permutations)

	return permutations
}

func dfsPermute(nums, solution []int, visited map[int]bool, permutations *[][]int) {
	for index, value := range nums {
		// 如果当前节点没有访问 将当前节点加入解决方案
		if exists, isVisited := visited[index]; !exists || !isVisited {
			visited[index] = true
			solution = append(solution, value)
			length := len(solution)
			// 解决方案长度和数组源长度一样 说明已经遍历完成
			if length == len(nums) {
				*permutations = append(*permutations, solution)
			} else {
				s := make([]int, length)
				copy(s, solution)
				dfsPermute(nums, s, visited, permutations)
			}
			visited[index] = false
			solution = solution[:length-1]
		}
	}
}
