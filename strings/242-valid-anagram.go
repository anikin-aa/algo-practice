package main

// problem: https://leetcode.com/problems/valid-anagram/
// O(N) - time, O(1) - space
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	slice := make([]int, 26)

	for i := 0; i < len(s); i++ {
		slice[s[i]-'a'] += 1
		slice[t[i]-'a'] -= 1
	}

	for _, x := range slice {
		if x != 0 {
			return false
		}
	}
	return true
}
