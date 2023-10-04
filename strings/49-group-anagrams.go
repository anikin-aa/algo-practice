package main

import "sort"

// https://leetcode.com/problems/group-anagrams/description/
// O(N * Klog(K)) - time
// O(N) - space
// K - max lenght of string from strs, N - len of strs
func groupAnagrams(strs []string) [][]string {
	grams := map[string][]string{}

	for _, str := range strs {
		sorted := sortString(str)
		if _, ok := grams[sorted]; !ok {
			grams[sorted] = append(grams[sorted])
		} else {
			grams[sorted] = append(grams[sorted], str)
		}
	}

	res := make([][]string, 0, len(grams))

	for _, groups := range grams {
		res = append(res, groups)
	}

	return res
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
