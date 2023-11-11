package bits

/*
https://leetcode.com/problems/counting-bits/
*/
func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		res[i] = getBits(i)
	}

	return res
}

func getBits(n int) int {
	res := 0

	for n != 0 {
		res += 1
		n = n & (n - 1)
	}

	return res
}
