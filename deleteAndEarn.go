package main

// https://leetcode.com/problems/delete-and-earn/

func max2(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func deleteAndEarn(nums []int) int {
	gains := map[int]int{}
	max := 0
	for i := range nums {
		if _, ok := gains[nums[i]]; ok {
			gains[nums[i]] += nums[i]
		} else {
			gains[nums[i]] = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	calc := make([]int, max+1)
	add := 0
	if _, ok := gains[1]; ok {
		add = gains[1]
	}
	calc[1] = add

	for i := 2; i <= max; i++ {
		gain := 0
		if _, ok := gains[i]; ok {
			gain = gains[i]
		}
		calc[i] = max2(calc[i-1], calc[i-2]+gain)
	}

	return calc[max]
}
