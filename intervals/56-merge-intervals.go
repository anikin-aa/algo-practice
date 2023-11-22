package intervals

import "sort"

// https://leetcode.com/problems/merge-intervals/
// time: O(n) + O(nlogn)
// space: O(n)

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    res := [][]int{
        intervals[0],
    }

    for i := 1; i < len(intervals); i++ {

        last, cur := res[len(res) - 1], intervals[i]

        if last[1] >= cur[0] && last[1] <= cur[1] {
            res[len(res) - 1][1] = cur[1]
        } else if cur[1] > last[1] {
            res = append(res, cur)
        }

    }

    return res
}