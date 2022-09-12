package main

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

// s = "asfdl;kajwefo;ajfea;oweifja;fowefijaweofjiaowef"
// sub3 = fdl
// sub4 = dl;

// for each window size
//    compute unique for substring
//    for each substring
//      update unique
//        subtract first char
//        add last char

func uniqueLetterString(s string) int {
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
