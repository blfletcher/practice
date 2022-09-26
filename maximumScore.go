package main

// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations

func max3(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func calcScore_td(i, left int, multipliers, nums []int, memo [][]int) int {
	if i == len(multipliers) {
		return 0
	}

	if memo[i][left] == 0 {
		right := len(nums) - 1 - (i - left)
		memo[i][left] = max3(multipliers[i]*nums[left]+calcScore_td(i+1, left+1, multipliers, nums, memo), multipliers[i]*nums[right]+calcScore_td(i+1, left, multipliers, nums, memo))
	}

	return memo[i][left]
}

// top down
func maximumScore_td(nums []int, multipliers []int) int {
	memo := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, len(multipliers))
	}

	return calcScore_td(0, 0, multipliers, nums, memo)
}

// bottom up
func maximumScore(nums []int, multipliers []int) int {
	n := len(nums)
	m := len(multipliers)

	tab := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		tab[i] = make([]int, m+1)
	}

	for i := len(multipliers) - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			mult := multipliers[i]
			right := n - 1 - (i - left)
			tab[i][left] = max3(mult*nums[left]+tab[i+1][left+1], mult*nums[right]+tab[i+1][left])
		}
	}

	return tab[0][0]
}
