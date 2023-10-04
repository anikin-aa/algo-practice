package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// O(n) - time, O(1) - space
func lengthOfLongestSubstring(s string) int {
	l, r, res := 0, 0, 0
	window := map[byte]bool{}

	for r < len(s) {
		if _, ok := window[s[r]]; !ok {
			window[s[r]] = true
			res = max(res, len(window))
			r += 1
		} else {
			delete(window, s[l])
			l += 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
