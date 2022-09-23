package main

import (
	"math"
)

// https://leetcode.com/problems/sliding-window-maximum/

func maxSlidingWindow(nums []int, k int) []int {
	results := make([]int, len(nums)-k+1)
	queue := []int{}

	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		if i >= k-1 {
			results[i-k+1] = nums[queue[0]]
		}

		if queue[0] == i-k+1 {
			queue = queue[1:]
		}
	}

	return results
}

func findMax(idx, k int, nums []int) (int, int) {
	maxIdx, maxValue := 0, math.MinInt

	for i := idx; i < idx+k; i++ {
		if nums[i] > maxValue {
			maxIdx = i
			maxValue = nums[i]
		}
	}

	return maxIdx, maxValue
}

func maxSlidingWindow_v1(nums []int, k int) []int {
	results := make([]int, len(nums)-k+1)
	maxIdx, maxValue := findMax(0, k, nums)
	results[0] = maxValue

	for i := 1; i <= len(nums)-k; i++ {
		if i > maxIdx {
			maxIdx, maxValue = findMax(i, k, nums)
		} else if nums[i+k-1] > maxValue {
			maxIdx = i + k - 1
			maxValue = nums[maxIdx]
		}
		results[i] = maxValue
	}

	return results
}
