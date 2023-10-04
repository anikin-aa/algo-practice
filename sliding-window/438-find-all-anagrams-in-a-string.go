package main

// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
// O(n + m) - time
// O(1) - seems like the space is constant (becase alphabet is finite), but not sure
func findAnagrams(s string, p string) []int {
	res := []int{}
	k := len(p)
	src, cur := [26]int{}, [26]int{}
	wSet := map[[26]int]bool{}

	for i := 0; i < k; i++ {
		src[p[i]-'a'] += 1
	}

	wSet[src] = true

	l, r := 0, 0

	for r < len(s) {
		cur[s[r]-'a'] += 1

		if r >= k {
			cur[s[l]-'a'] -= 1
			l += 1
		}

		if _, ok := wSet[cur]; ok {
			res = append(res, l)
		}

		r += 1
	}

	return res
}

// TLE
// func findAnagrams(s string, p string) []int {
//     res := []int{}
//     // window size
//     k := len(p)
//     // p = sortString(p)

//     for i := 0; i < len(s) - k + 1; i++ {
//         if p == sortString(s[i:i+k]) {
//             res = append(res, i)
//         }
//     }

//     return res
// }

// func sortString(s string) string {
//     r := []rune(s)
//     sort.Slice(r, func(i, j int) bool {
//         return r[i] < r[j]
//     })
//     return string(r)
// }
