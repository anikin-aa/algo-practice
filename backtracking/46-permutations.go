package backtracking

// https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
    res := [][]int{}
    n := len(nums)
    cur := []int{}
    visit := make([]bool, n)

    var dfs func()

    dfs = func() {
        
        if len(cur) == n {
            cp := make([]int, len(cur))
            copy(cp, cur)
            res = append(res, cp)
            return
        }

        for i := 0; i < n; i++ {
            if visit[i] {continue}
            visit[i] = true
            cur = append(cur, nums[i])
            dfs()
            cur = cur[:len(cur) - 1]
            visit[i] = false
        }

    }

    dfs()

    return res   
}
}