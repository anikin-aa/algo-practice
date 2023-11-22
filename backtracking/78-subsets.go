package backtracking

// https://leetcode.com/problems/subsets/
func subsets(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}

	var dfs func(idx int)

	dfs = func(idx int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		res = append(res, tmp)

		for i := idx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0)

	return res
}
