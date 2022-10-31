package main

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	charMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	all := []string{}
	for i := range digits {
		digit := digits[i]

		new := []string{}
		chars := charMap[digit]
		for j := range chars {
			char := chars[j]

			if i > 0 {
				for k := range all {
					new = append(new, all[k]+string(char))
				}
			} else {
				new = append(new, string(char))
			}
		}
		all = new
	}

	return all
}
