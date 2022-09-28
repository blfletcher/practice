package main

// https://leetcode.com/problems/candy/

func candies(i int, memo, ratings []int) int {
	if memo[i] == 0 {
		left, right := 1, 1
		if i > 0 && ratings[i] > ratings[i-1] {
			left = candies(i-1, memo, ratings) + 1
		}
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right = candies(i+1, memo, ratings) + 1
		}

		max := left
		if right > left {
			max = right
		}

		memo[i] = max
	}

	return memo[i]
}

func candy(ratings []int) int {
	memo := make([]int, len(ratings))

	total := 0
	for i := range ratings {
		total += candies(i, memo, ratings)
	}

	return total
}

/*func candies(i int, memo, ratings []int) (int, []int) {
	if memo[i] == 0 {
		if i == 0 {
			memo[i] = 1
		} else if ratings[i] > ratings[i-1] {
			memo[i], memo = candies(i-1, memo, ratings)
			memo[i]++
		} else {
			memo[i] = 1
		}
	}

	return memo[i], memo
}*/
