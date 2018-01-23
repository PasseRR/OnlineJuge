package page1

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	// 最终结果
	result := [][]int{}
	// 初始化结果集
	combination := [][]int{}
	sort.Ints(candidates)
	backtrackingSum2(candidates, []int{}, target, &combination)
	// 去重
	for _, sub := range combination {
		sort.Ints(sub)
		if !isSolutionExists(sub, result) {
			result = append(result, sub)
		}
	}
	return result
}

func isSolutionExists(solution []int, result [][]int) bool {
	for _, sub := range result {
		if len(sub) != len(solution) {
			continue
		}

		flag := true
		for j, value := range sub {
			if solution[j] != value {
				flag = false
				break
			}
		}

		if flag {
			return flag
		}
	}

	return false
}

func backtrackingSum2(candidates, solution []int, target int, combination *[][]int) {
	for i, candidate := range candidates {
		if candidate < target {
			solution = append(solution, candidate)
			// 每个元素只能使用一次 复制candidates并移除当前元素
			c := make([]int, len(candidates))
			copy(c, candidates)
			// 减去已经加上的数字
			backtrackingSum2(append(c[:i], c[i+1:]...), solution, target-candidate, combination)
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
