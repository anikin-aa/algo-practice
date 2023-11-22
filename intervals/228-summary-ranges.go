package intervals

import "fmt"

// https://leetcode.com/problems/summary-ranges/
// time: O(n)
// space: O(n)
func summaryRanges(nums []int) []string {
	res := []string{}

	if len(nums) == 0 {
		return res
	}

	// 0,1,2,4,5,7

	// 0,1,2,3,4,5

	start := nums[0]

	// i = 1; nums[0] (0) + 1 == nums[1](1), continue
	// i = 2; nums[1] (1) + 1 == nums[2](2), continue
	// i = 3; nums[2] (2) + 1 == nums[3](4), else, start = 0, nums[2] = 2 != , res = append(res, 0->2) start = 4
	
	// i = 4; nums[3] (4) + 1 == nums[4] (5), continue
	// i = 5; nums[4] (5) + 1 == nums[5] (7), else start = 4, nums[4] = 5 !=, res = append(res, 4->5) start = 7

	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 == nums[i] {
			continue
		} else {
			str := fmt.Sprintf("%d->%d", start, nums[i-1])
			if start == nums[i-1] {
				str = fmt.Sprintf("%d", start)
			}
			res = append(res, str)
			start = nums[i]
		}
	}

	str := fmt.Sprintf("%d->%d", start, nums[len(nums) - 1])
	if start == nums[len(nums) - 1] {
		str = fmt.Sprintf("%d", start)
	}

	res = append(res, str)

	return res
}
