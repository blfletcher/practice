package main

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

func longestSubarray(nums []int, limit int) int {
	currStart := 0
	maxSize := 0
	maxQueue, minQueue := []int{}, []int{}

	for i := range nums {
		num := nums[i]

		for len(maxQueue) > 0 {
			idx := maxQueue[len(maxQueue)-1]
			if nums[idx] <= num {
				maxQueue = maxQueue[:len(maxQueue)-1]
			} else {
				break
			}
		}
		for len(minQueue) > 0 {
			idx := minQueue[len(minQueue)-1]
			if nums[idx] > num {
				minQueue = minQueue[:len(minQueue)-1]
			} else {
				break
			}
		}
		maxQueue = append(maxQueue, i)
		minQueue = append(minQueue, i)
		diff := nums[maxQueue[0]] - nums[minQueue[0]]
		if diff <= limit {
			size := i - currStart + 1
			if size > maxSize {
				maxSize = size
			}
		} else {
			currStart++
			if currStart > maxQueue[0] {
				maxQueue = maxQueue[1:]
			}
			if currStart > minQueue[0] {
				minQueue = minQueue[1:]
			}
		}
	}

	return maxSize
}
