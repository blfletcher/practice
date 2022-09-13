package main

import "fmt"

// https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/
// power set == all subsets

// 000
// 001
// 010
// 011
// 1<<n == 2^n
// n > ~16, this is bad
func iteratePowerset(things []string) []string {
	n := len(things)
	subsets := []string{}
	for mask := 0; mask < 1<<n; mask++ {
		value := ""
		for j := 0; j < n; j++ {
			if (mask & (1 << j)) == 1 {
				value += things[j]
			}
		}
		subsets = append(subsets, value)
	}

	return subsets
}

// idx + c1
// 0 idx + idx - prev
// prev      idx
// stakefwfweabcdea
//   ^       ^    ^
// sum
// new prefixes
// a_final = idx - prev
// prev_c1 = old prefixes
// overlap_total = old prefixes * new prefixes (suffixes)
// sum += a_final + overlap_total
// prev = idx
// prev_c1 = a_final

// sta, ta, a, stak, tak, ak, ka, kab, ab

// AB
// idx = 1
// c1 = 1
// over = 0
// previdx[A] = 1
// cs[A] = 1
// total = 2

// "C"
// abCabdCefghC = 37
// ABBA == A, AB, B, B, BBA, BA, A
func uniqueSingleLetterString(s string) int {
	prev := -1
	prevC1 := 0
	total := 1
	for i := 0; i < len(s); i++ {
		idx := i
		// B == 1
		// B == 2
		// A == 3
		c1 := idx - prev
		// A = 1
		// B = 2
		// B = 1
		// A == 2
		total += prevC1*(c1-1) + c1
		// A == 1
		// B == 3
		// B == 4
		// A == 7
		prev = idx
		// A == 0
		// B == 1
		// B == 2
		// A == 3
		prevC1 = c1
		// A == 1
		// B == 2
		// B == 1
		// A == 2
	}

	return total
}

// ABC
func uniqueLetterString(s string) int {
	cs := make([]int, 26)
	prevIdxs := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	n := int(s[0] - 'A')
	cs[n] = 1
	prevIdxs[n] = 0

	total := 1
	for i := 1; i < len(s); i++ {
		idx := i
		char := int(s[i] - 'A')

		c1 := idx - prevIdxs[char]
		overlapC := cs[char] * (c1 - 1)

		total += overlapC + c1
		cs[char] = c1
		prevIdxs[char] = idx
		fmt.Printf("\n\nidx [%d] char [%c] charIdx [%d] c1 [%d] overlap [%d] total [%d]", idx, s[i], char, c1, overlapC, total)
	}

	return total
}

// prev += c1 - cs[char]
// prev = prev + idx - prevIdxs[char] - cs[char]
// prev = prev + i - index[v] - count[v]

func old_uniqueLetterString(s string) int {
	uniqueChars := 0
	// Iterate over lengths
	for i := 1; i <= len(s); i++ {
		checkUnique := make([]int, 26)
		uniqueSubChars := 0

		// initialize for unique
		for k := 0; k < i; k++ {
			chr := s[k] - 'A'
			checkUnique[chr] += 1
			if checkUnique[chr] == 1 {
				uniqueSubChars++
			} else if checkUnique[chr] == 2 {
				uniqueSubChars--
			}
			uniqueChars += uniqueSubChars
		}

		// Shift window of size i
		for j := 1; j < len(s)-i; j++ {
			chr1 := s[j+i] - 'A'
			chr2 := s[j-1] - 'A'
			checkUnique[chr1] += 1
			checkUnique[chr2] -= 1
			// add uniqueness of s[j+i]
			if checkUnique[chr1] == 1 {
				uniqueSubChars++
			} else if checkUnique[chr1] == 2 {
				uniqueSubChars--
			}
			// subtract uniqueness of s[j-1]
			if checkUnique[chr2] == 1 {
				uniqueSubChars++
			} else if checkUnique[chr2] == 0 {
				uniqueSubChars--
			}
			uniqueChars += uniqueSubChars
		}
	}

	return uniqueChars
}
