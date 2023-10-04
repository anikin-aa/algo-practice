package main

// https://leetcode.com/problems/sliding-window-maximum/description/
// Time - O(n), space - O(k)

type deque struct {
	h []int
}

func (d *deque) push(i int) {
	d.h = append(d.h, i)
}

func (d *deque) getFirst() int {
	return d.h[0]
}

func (d *deque) getLast() int {
	return d.h[len(d.h)-1]
}

func (d *deque) popLast() {
	d.h = d.h[:len(d.h)-1]
}

func (d *deque) popFirst() {
	d.h = d.h[1:]
}

func (d *deque) empty() bool {
	return len(d.h) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return []int{}
	} else if k == 1 {
		return nums
	}

	res := []int{}
	d := &deque{}

	for i := 0; i < len(nums); i++ {

		//
		if !d.empty() && d.getFirst() == i-k {
			d.popFirst()
		}

		for !d.empty() && nums[d.getLast()] < nums[i] {
			d.popLast()
		}

		d.push(i)

		if i >= k-1 {
			res = append(res, nums[d.getFirst()])
		}

	}

	return res
}

/* TLE

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}

	for i := 0; i < len(nums) - k + 1; i ++ {
		res = append(res, getMax(nums[i:i+k]))
	}

	return res
}

func getMax(nums []int) int {
	res := math.MinInt

	for _, x := range nums {
		res = max(x, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {return a}
	return b
}
*/
