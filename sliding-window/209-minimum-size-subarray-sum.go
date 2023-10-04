package main

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum/
// O(n) - time, O(1) - space
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	curSum := 0
	l, r := 0, 0

	for r < len(nums) {

		curSum += nums[r]

		for curSum >= target {
			curSum -= nums[l]
			res = min(res, r-l+1)
			l += 1
		}
		r += 1
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
ex 1

2,3,1,2,4,3

l
      r

sum = 8
*/
