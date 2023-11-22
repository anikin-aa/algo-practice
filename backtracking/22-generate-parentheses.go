package backtracking

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
    res := []string{}

    cur := []byte{}

    var dfs func(int, int)

    dfs = func(o, c int) {
        if len(cur) == 2 * n {
            res = append(res, string(cur))
            return
        }

        if o < n {
            cur = append(cur, '(')
            o += 1
            dfs(o, c)
            o -= 1
            cur = cur[:len(cur) - 1]
        }

        if c < o {
            cur = append(cur, ')')
            c += 1
            dfs(o, c)
            c -= 1
            cur = cur[:len(cur) - 1]
        }
    }

    dfs(0, 0)

    return res
}
