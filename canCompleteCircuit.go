package main

// https://leetcode.com/problems/gas-station/

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	startIdx := 0
	current := 0
	total := 0
	for i := 0; i < n; i++ {
		total += gas[i] - cost[i]
		current += gas[i] - cost[i]
		if current < 0 {
			startIdx = i + 1
			current = 0
		}
	}

	res := -1
	if startIdx < n && total >= 0 {
		res = startIdx
	}
	return res
}
