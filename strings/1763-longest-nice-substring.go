package main

import "unicode"

// https://leetcode.com/problems/longest-nice-substring/
// O(n^3) - time, O(1) - space
func longestNiceSubstring(s string) string {
	res := ""

	nice := func(s string) bool {
		// save presence of upper and lower case of runes
		lower, upper := make([]int, 26), make([]int, 26)

		for _, r := range s {
			if unicode.IsUpper(r) {
				upper[r-'A'] += 1
			} else {
				lower[r-'a'] += 1
			}
		}

		for idx := range lower {
			if lower[idx] != 0 {
				if upper[idx] == 0 {
					return false
				}
			}
		}

		return true
	}

	// bruteforce over all of the substrings from s
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if !nice(s[i:j]) {
				continue
			}
			if len(s[i:j]) > len(res) {
				res = s[i:j]
			}
		}
	}

	return res
}
