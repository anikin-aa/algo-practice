package main

/*
O(n) space, O(1) time
*/
type NumArray303 struct {
	sum []int
}

func Constructor(nums []int) NumArray303 {
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	return NumArray303{sum}
}

func (n *NumArray303) SumRange(left int, right int) int {
	return n.sum[right+1] - n.sum[left]
}
