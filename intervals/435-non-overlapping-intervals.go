package intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	// res will contain non-overlapping
	res := [][]int{}


	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})





	return len(intervals) - len(res)
}