package page1

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// 最终结果
	result := [][]int{}
	// 初始化结果集
	combination := [][]int{}
	// 排序、去重元数据
	sort.Ints(candidates)

	backtrackingSum(candidates, []int{}, target, &combination)
	// 去重
	for _, sub := range combination {
		length := len(sub)
		if length == 1 {
			result = append(result, sub)
		} else {
			flag := true
			for i := 1; i < length; i++ {
				if sub[i-1] > sub[i] {
					flag = false
					break
				}
			}
			if flag {
				result = append(result, sub)
			}
		}
	}

	return result
}

func backtrackingSum(candidates, solution []int, target int, combination *[][]int) {
	for _, candidate := range candidates {
		if candidate < target {
			solution = append(solution, candidate)
			// 减去已经加上的数字
			backtrackingSum(candidates, solution, target-candidate, combination)
			// 回溯
			solution = solution[:len(solution)-1]
		} else {
			if candidate == target {
				solution = append(solution, candidate)
				// 复制slice
				s := make([]int, len(solution))
				copy(s, solution)
				*combination = append(*combination, s)
			}

			break
		}
	}
}
