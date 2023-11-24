package intervals

import "sort"

// https://leetcode.com/problems/non-overlapping-intervals/description/
// time - O(nlogn)
// space - O(n)
func eraseOverlapIntervals(intervals [][]int) int {

    sort.Slice(intervals, func(i, j int) bool {
        a, b := intervals[i], intervals[j]
        // sort by ends, small first
        return a[1] < b[1]
    })

    nonOverlap := [][]int{
        intervals[0],
    }

    for i := 1; i < len(intervals); i++ {
        cur, prev := intervals[i], nonOverlap[len(nonOverlap) - 1]

        // non-overlapping
        if prev[1] <= cur[0] {
            nonOverlap = append(nonOverlap, cur)
        }
    }

    return len(intervals) - len(nonOverlap)
}