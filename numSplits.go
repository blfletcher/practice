package main

// https://leetcode.com/problems/number-of-good-ways-to-split-a-string/

// good split == split s into 2 non empty strings (sleft, sright)
//   sleft + sright == s
//   unique chars(sleft) == unique chars(sright)

// what to know
//   hash map[character] = 2D array of first and last character positions
//   [len(s)][len(hash map)]
//   iterate string
//     iterate hash map, counting characters <= index, end > index
//     if ==, save result

// aacaba
// a = 0, 5
// c = 2, 2
// b = 4, 4
// [ 1, 0, 2, 0, 3, 3]

func numSplits_ben(s string) int {
	chars := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := chars[s[i]]; ok {
			chars[s[i]][1] = i
		} else {
			chars[s[i]] = []int{i, i}
		}
	}

	good := 0
	for i := 0; i < len(s); i++ {
		left, right := 0, 0
		for c := range chars {
			char := chars[c]
			if char[0] <= i {
				left++
			}
			if char[1] > i {
				right++
			}
		}
		if left == right {
			good++
		}
	}

	return good
}

const incLeft = 1
const incLeftDecRight = 2
const decRight = 3

func numSplits(s string) int {
	chars := map[byte][]int{}
	// Map characters to first, last positions
	for i := 0; i < len(s); i++ {
		if _, ok := chars[s[i]]; ok {
			chars[s[i]][1] = i
		} else {
			chars[s[i]] = []int{i, i}
		}
	}

	// Serialize map to state vector
	state := make([]int, len(s))
	for c := range chars {
		char := chars[c]
		if char[0] == char[1] {
			state[char[0]] = incLeftDecRight
		} else {
			state[char[0]] = incLeft
			state[char[1]] = decRight
		}
	}

	// Count unique characters in state vector
	left := 0
	right := len(chars)
	good := 0
	for i := 0; i < len(s); i++ {
		switch state[i] {
		case incLeft:
			left++
		case incLeftDecRight:
			left++
			right--
		case decRight:
			right--
		}
		if left == right {
			good++
		}
	}

	return good
}
