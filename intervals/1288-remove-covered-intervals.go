package intervals

import "sort"

// https://leetcode.com/problems/remove-covered-intervals/description/
// time: O(n) + O(nlogn)
// space: O(n) because of the res
func removeCoveredIntervals(intervals [][]int) int {
    // 1, 4; 2, 8; 3,6

    sort.Slice(intervals, func(i, j int) bool {
        // if start of intervals are equal, intervals with largest end should be first 
        // [1,2]; [1,4]; [3,4]
        // will be sorted to [1,4]; [1,2]; [3,4]
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    res := [][]int{intervals[0]}

    for i := 1; i < len(intervals); i++ {
        prev, cur := res[len(res) - 1], intervals[i]
        a, b, c, d := cur[0], cur[1], prev[0], prev[1]

        if c <= a && b <= d {
            continue
        } else {
            res = append(res, cur)
        }

    }


    return len(res)
}