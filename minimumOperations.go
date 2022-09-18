package main

import (
	"sort"
)

// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/

func minimumOperations(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	x := 0
	operations := 0

	for _, v := range nums {
		val := v

		diff := val - x
		if val > 0 && diff > 0 {
			operations++
			if diff > 0 {
				x = val
			}
		}
	}

	return operations
}
