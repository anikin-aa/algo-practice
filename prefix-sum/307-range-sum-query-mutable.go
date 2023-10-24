package main

// 1st non-optimal approach
// recalculate on each update
type NumArray307 struct {
	holder []int
	sum    []int
}

func Constructor(nums []int) NumArray307 {
	n := len(nums)
	sum := make([]int, n+1)

	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	return NumArray307{holder: nums, sum: sum}
}

func (n *NumArray307) Update(index int, val int) {
	n.holder[index] = val
	for i := index + 1; i < len(n.sum); i++ {
		n.sum[i] = n.holder[i-1] + n.sum[i-1]
	}
}

func (n *NumArray307) SumRange(l int, r int) int {
	return n.sum[r+1] - n.sum[l]
}

// 2nd approach with segment tree
