package main


/*
O(n) - space, O(n^2) - time
*/
func subarraySum(nums []int, k int) int {
	prefixSum := make([]int, len(nums) + 1)
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i - 1] + nums[i - 1]
	}
	res := 0

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j + 1] - nums[i] == k {res += 1}
		}
	}

	return res
}


/*
O(n) - space, O(n) - time
*/
func subarraySum2(nums []int, k int) int {
    prefixSum := map[int]int{0: 1}
    res, sum := 0, 0 

    for _, x := range nums {
        sum += x
        res += prefixSum[sum - k]
        prefixSum[sum] += 1
    }

    return res
}