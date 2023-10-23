package main

/*
O(n) - time, O(1) - space 
*/
func pivotIndex(nums []int) int {
    totalSum := 0

    for _, x := range nums {totalSum += x}

    // 1,7,3,6,5,6

    // totalSum = 28

    curSum := 0

    for i, x := range nums {
        if curSum == totalSum - curSum - x {
            return i
        }
        curSum += x
    }

    return -1
}