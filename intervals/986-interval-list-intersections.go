package intervals


// https://leetcode.com/problems/interval-list-intersections/
// 
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {


	res := [][]int{}

	i, j := 0, 0

	// [0,2],[5,10],[13,23],[24,25]
    // [1,5],[8,12],[15,24],[25,26]
	// i, j = 0, 0; start, end = 1, 2; start <= end ? add res (1,2); 2 <= 5 ? i += 1
	// i, j = 1, 0; start, end = 5, 5; start <= end ? add res (5,5); 10 <= 5 ? j += 1
	// i, j = 1, 1; start, end = 8, 10; start <= end ? add res (8,10); 10 <= 12 ? i += 1
	// i, j = 2, 1; start, end = 13, 12; start <= end ? ; 23 <= 12 ? j += 1
	// i, j = 2, 2; start, end = 15, 23; start <= end ? add res (15, 23); 23 <= 24 ? i += 1
	// i, j = 3, 2; start, end = 24, 24; start <= end ? add res (24, 24); 25 <= 24 ? j += 1
	// i, j = 3, 3; start, end = 25, 25; start <= end ? add res (25, 25); 25 <= 25 ? j += 1

	// [1,3],[5,9]
	// [4,4],[10,11]

	// i, j = 0,0; start, end = 4,3; start <= end ? ; 3 <= 4 ? i += 1
	// i, j = 1,0; start, end = 5,9; start <= end ? ; 9 <= 4 ? j += 1
	// i, j = 1,1; start, end = 10,9; start <= end ? ; 9 <= 11 ? i += 1


	for i < len(firstList) && j < len(secondList) {
		a, b := firstList[i], secondList[j]

		start, end := max(a[0], b[0]), min(a[1], b[1])

		if start <= end {
			res = append(res, []int{start, end})
		}

		if a[1] <= b[1] {
			i += 1
		} else {
			j += 1
		}

	}


	return res
}

func min(a, b int) int {
	if a < b {return a}
	return b
}

func max(a, b int) int {
	if a > b {return a}
	return b
}