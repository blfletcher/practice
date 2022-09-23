package main

// https://leetcode.com/problems/substring-with-largest-variance/

/*
	iterate string, finding characters that are present
	double iteration of present characters
	for i character present
		for j character present
			if i == j, skip
			else, find max of the 2 characters

	find max(lowFrequency, highFrequency):
		result = 0
		lowFrequency, highFrequency count = 0
		iterate string
			if lowFrequency
				increment lowFrequency count
			else if highFrequency
				increment highFrequency count
			else
				skip to next character
			if lowFrequency count is 0
				result = max(result, highFrequency count - 1)
			else result = 0
				result = max(result, highFrequency count - lowFrequency count)
			if highFrequency count - lowFrequency count < 0
				set lowFrequency,highFrequency count to 0
*/

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func kadaneMax(s string, currentMax int, lowFrequencyChar, highFrequencyChar byte) int {
	lowCount, highCount := 0, 0
	for i := range s {
		c := s[i]

		if c == lowFrequencyChar {
			lowCount++
		} else if c == highFrequencyChar {
			highCount++
		} else {
			continue
		}

		if lowCount == 0 {
			currentMax = max(currentMax, highCount-1)
		} else {
			currentMax = max(currentMax, highCount-lowCount)
		}

		if highCount < lowCount {
			lowCount, highCount = 0, 0
		}
	}

	return currentMax
}

func largestVariance(s string) int {
	present := map[byte]bool{}
	for i := range s {
		c := s[i]

		if _, ok := present[c]; !ok {
			present[c] = true
		}
	}

	max := 0
	for i := range present {
		for j := range present {
			if i == j {
				continue
			}
			max = kadaneMax(s, max, i, j)
		}
	}

	return max
}

// iterate the string
// finding local maximum, maximum of all previous substrings
// local maximum = current value + maximum of all previous substrings

// counter := [][]int{{-1,-1,-1} * 26}
// beginning, end, num occurrences
// iterate string to populate counter
// iterate counter

// bab
// max = 0
// iterate string i
//   local max = make([]int, 26), initialized to -1, except s[i] set to 1
//
//   iterate j = i; j < len(s)
//

// babbabba
// for i := range s {
// b = 0
// [{-1,-1,-1}, {0,0,1}]
// b, ba, a = 0
// [{1,1,1},{0,0,1}]
// bab = 0
// b, ba, a
// [{1,1,1},{0,2,2}]
// babb
// [{1,1,1},{0,3,3}]
// babba
// [{1,4,2},{0,3,3}]
// babbab
// [{1,4,2},{0,5,4}]
// babbabb
// [{1,4,2},{0,6,5}]
// babbabba
// [{1,7,3},{0,6,5}]

// babbcbba
// for i := range s {
// b
// [{-1,-1,-1},{0,0,1},{-1,-1,-1}]
// ba
// [{1,1,1},{0,0,1},{-1,-1,-1}]
// bab
// [{1,1,1},{0,2,2},{-1,-1,-1}]
// babb
// [{1,1,1},{0,3,3},{-1,-1,-1}]
// babbc
// [{1,1,1},{0,3,3},{4,4,1}]
// babbab
// [{1,1,1},{0,5,4},{4,4,1}]
// babbabb
// [{1,1,1},{0,6,5},{4,4,1}]
// babbabba
// [{1,7,2},{0,6,5},{4,4,1}]
//}

// init a character counting array[26] with 0
// maybe track min character, max character
// iterate string using i = 0 to len(str)
//   idx = str[i] - 'a'
//   array[idx]++
//   if array[idx] > max, set to max
//   if array[idx] < min, set to min
/*counter := make([]int, 26)
for i := range s {
	ii := i

	counter[s[ii]-'a']++
}

max, min := 0, len(s)
for i := range counter {
	ii := i

	if counter[ii] > max {
		max = counter[ii]
	}
	if counter[ii] < min && counter[ii] > 0 {
		min = counter[ii]
	}
}

return max - min*/
