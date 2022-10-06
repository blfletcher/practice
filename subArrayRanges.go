package main

// https://leetcode.com/problems/sum-of-subarray-ranges/

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	maxSum, minSum := int64(0), int64(0)
	maxQueue, minQueue := []int{}, []int{}

	for next := 0; next <= n; next++ {
		prevMax := -1
		for len(maxQueue) > 0 {
			i := maxQueue[len(maxQueue)-1]
			if next == n || nums[i] < nums[next] {
				maxQueue = maxQueue[:len(maxQueue)-1]
				if len(maxQueue) > 0 {
					prevMax = maxQueue[len(maxQueue)-1]
				} else {
					prevMax = -1
				}
				maxSum += int64(nums[i]) * int64(next-i) * int64(i-prevMax)
			} else {
				break
			}
		}

		prevMin := -1
		for len(minQueue) > 0 {
			i := minQueue[len(minQueue)-1]
			if next == n || nums[i] > nums[next] {
				minQueue = minQueue[:len(minQueue)-1]
				if len(minQueue) > 0 {
					prevMin = minQueue[len(minQueue)-1]
				} else {
					prevMin = -1
				}
				minSum += int64(nums[i]) * int64(next-i) * int64(i-prevMin)
			} else {
				break
			}
		}
		maxQueue = append(maxQueue, next)
		minQueue = append(minQueue, next)
	}

	return maxSum - minSum
}
