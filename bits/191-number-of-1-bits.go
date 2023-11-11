package bits

/*
https://leetcode.com/problems/number-of-1-bits/
*/

func hammingWeight(n uint32) int {
	res := 0

	for n != 0 {
		res += 1

		n = n & (n - 1)
	}

	return res
}
