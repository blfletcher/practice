package main

// https://leetcode.com/problems/trapping-rain-water/

// DP
func trap(height []int) int {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		max := height[i]
		if max < leftMax[i-1] {
			max = leftMax[i-1]
		}

		leftMax[i] = max
	}
	for i := n - 2; i >= 0; i-- {
		max := height[i]
		if max < rightMax[i+1] {
			max = rightMax[i+1]
		}

		rightMax[i] = max
	}

	totalRain := 0
	for i := 1; i < n-1; i++ {
		min := leftMax[i]
		if min > rightMax[i] {
			min = rightMax[i]
		}
		totalRain += min - height[i]
	}

	return totalRain
}

// Stack
func trap_stack(height []int) int {
	totalRain := 0
	stack := []int{0}
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				dist := i - top - 1
				hole := height[i]
				if hole > height[top] {
					hole = height[top]
				}
				totalRain += dist * (hole - height[prev])
			}
		}
		stack = append(stack, i)
	}

	return totalRain
}
