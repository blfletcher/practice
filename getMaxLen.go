package main

// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/

func getMaxLen(nums []int) int {
	nums = append(nums, 0)

	maxLen := 0
	currStart := 0
	currFirstNeg := -1
	currLastNeg := -1
	negProduct := 1
	for i := range nums {
		num := nums[i]

		if num < 0 {
			negProduct *= -1
			if currFirstNeg < 0 {
				currFirstNeg = i + 1
			}
			currLastNeg = i
		} else if num == 0 {
			currLen := i - currStart
			if negProduct < 0 {
				currLen = currLastNeg - currStart
				if i-currFirstNeg > currLen {
					currLen = i - currFirstNeg
				}
			}
			if currLen > maxLen {
				maxLen = currLen
			}
			currStart = i + 1
			currFirstNeg = -1
			currLastNeg = -1
			negProduct = 1
		}
	}

	return maxLen
}
