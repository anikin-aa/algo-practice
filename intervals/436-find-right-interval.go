package intervals

import "sort"

//https://leetcode.com/problems/find-right-interval/

func findRightInterval(intervals [][]int) []int {
	res := make([]int, len(intervals))

	idxs := map[int]int{}

	// save indexes
	for i, intr := range intervals {
		idxs[intr[0]] = i
	}

	// sort by start of interval
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	getRightIndex := func(from int) int {
		end := intervals[from]

		for j := from; j < len(intervals); j++ {
			start := intervals[j]
			// intervals sorted by start, 
			// no need to minimize start of interval
			if start[0] >= end[1] {
				return j
			}
		}
		return -1
	}

	for i, intr := range intervals {
		srcIndex := idxs[intr[0]]

		// start to search from i till end
		resIndex := getRightIndex(i)

		if resIndex == -1 {
			res[srcIndex] = -1
		} else {
			res[srcIndex] = idxs[intervals[resIndex][0]]
		}
	}

	return res
}
