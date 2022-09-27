package main

// https://leetcode.com/problems/plates-between-candles/

// a b c d e f
// 0  1    2       3
// a a+b a+b+c a+b+c+d

// prefix[3] - prefix[2] = sum(a[2..3])
// * * | * * * | * *
// 0 1 2 3 4 5 6 7 8
// 3,5 - begin=6, end = 1
// 2,5
// 0 0 1 1 1 1 2 2 2
// 1 2 2 3 4 5 5 6 7
//
// 2,6 = 1
// iterate the string, calculating the number of candles to that point
// take range, start from either end, looking for transition (value doesn't change), take the difference
// begin, end
// if begin > end -- empty
// prefix[end] - prefix[begin] + 1 == number of candles
// end - begin == number candles + plates
// #plates = (end - begin) - (prefix[end] - prefix[begin] + 1)

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	// Compute prefix sum of plates
	cnt := 0
	prefixSum := make([]int, n)

	// 0 == prev, 1 == next
	prevCandles := make([]int, n)
	nextCandles := make([]int, n)

	prevCandle := -1
	nextCandle := -1
	for i := 0; i < n; i++ {
		j := n - i - 1
		if s[i] == '*' {
			cnt++
		} else {
			prevCandle = i
		}
		if s[j] == '|' {
			nextCandle = j
		}
		prevCandles[i] = prevCandle
		nextCandles[j] = nextCandle
		prefixSum[i] = cnt
	}

	result := make([]int, len(queries))
	// Calculate query results
	for i := range queries {
		query := queries[i]

		// Ignore orphaned plates
		leftNextCandle := nextCandles[query[0]]
		rightPrevCandle := prevCandles[query[1]]
		if leftNextCandle < rightPrevCandle && leftNextCandle > -1 {
			result[i] = prefixSum[rightPrevCandle] - prefixSum[leftNextCandle]
		}
	}

	return result
}
