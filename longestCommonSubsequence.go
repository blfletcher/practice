package main

// https://leetcode.com/problems/longest-common-subsequence/

func longestCommonSubsequence(text1 string, text2 string) int {
	return longestCommonSubsequence_dynamic(text1, text2)
}

func longestCommonSubsequence_dynamic(text1, text2 string) int {
	long := text1
	short := text2
	if len(text2) > len(text1) {
		long = text2
		short = text1
	}

	prev := make([]int, len(short)+1)
	current := make([]int, len(short)+1)
	for j := len(long) - 1; j >= 0; j-- {
		for i := len(short) - 1; i >= 0; i-- {
			if short[i] == long[j] {
				current[i] = 1 + prev[i+1]
			} else {
				max := prev[i]
				if current[i+1] > max {
					max = current[i+1]
				}
				current[i] = max
			}
		}
		temp := prev
		prev = current
		current = temp
	}

	return prev[0]
}

func longestCommonSubsequence_memo(text1, text2 string) int {
	memo := make([][]int, len(text1))

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	max, memo := memoSolve(0, 0, text1, text2, memo)
	return max
}

func memoSolve(i, j int, text1, text2 string, memo [][]int) (int, [][]int) {
	if i == len(text1) || j == len(text2) {
		return 0, memo
	}

	if memo[i][j] != -1 {
		return memo[i][j], memo
	}

	v1, v2 := 0, 0
	answer := 0
	if text1[i] == text2[j] {
		v1, memo = memoSolve(i+1, j+1, text1, text2, memo)
		answer = 1 + v1
	} else {
		v1, memo = memoSolve(i, j+1, text1, text2, memo)
		v2, memo = memoSolve(i+1, j, text1, text2, memo)
		answer = v1
		if v2 > v1 {
			answer = v2
		}
	}
	memo[i][j] = answer

	return memo[i][j], memo
}

/*
func longestCommonSubsequence(text1 string, text2 string) int {
	long := text1
	short := text2
	if len(text1) < len(text2) {
		short = text1
		long = text2
	}

	shortMap := map[byte]int{}
	for i := range short {
		c := short[i]

		if _, ok := shortMap[c]; ok {
			shortMap[c]++
		} else {
			shortMap[c] = 1
		}
	}

	cnt := 0
	shortIdx := 0
	for i := range long {
		fmt.Printf("\n[%c]\n", long[i])
		if _, ok := shortMap[long[i]]; !ok {
			fmt.Printf("not in map\n\n")
			continue
		} else if shortMap[long[i]] == 0 {
			fmt.Printf("value zero\n\n")
			continue
		}

		for shortIdx < len(short) {
			if short[shortIdx] == long[i] {
				cnt++
				shortMap[long[i]]--
				shortIdx++
				break
			}
			if c, ok2 := shortMap[short[shortIdx]]; ok2 && c > 0 {
				shortMap[short[shortIdx]]--
			}
			shortIdx++
		}
	}

	return cnt
}
*/
