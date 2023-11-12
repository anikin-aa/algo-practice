package backtracking

// https://leetcode.com/problems/combinations/
func combine(n int, k int) [][]int {
	// nums := make([]int, n)
	// for i := range nums {
	// nums[i] = i + 1
	// }
	res := [][]int{}

	cur := []int{}

	var dfs func(int)

	dfs = func(idx int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)

		if len(tmp) == k {
			res = append(res, tmp)
			return
		}

		// for i := idx; i < len(nums); i++ {
		// cur = append(cur, nums[i])

		for i := idx; i <= n; i++ {
			cur = append(cur, i)
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(1)

	return res
}
