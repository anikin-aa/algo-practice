package backtracking

import "sort"

// https://leetcode.com/problems/combination-sum/description/
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	cur := []int{}

	sort.Ints(candidates)

	var dfs func(int)

	dfs = func(idx int) {
		if target < 0 {
			return
		}
		if target == 0 {
			cp := make([]int, len(cur))
			copy(cp, cur)
			res = append(res, cp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			x := candidates[i]
			cur = append(cur, x)
			target -= x
			dfs(i)
			target += x
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0)

	return res
}
