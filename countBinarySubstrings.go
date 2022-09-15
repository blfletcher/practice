package main

// https://leetcode.com/problems/count-binary-substrings/

// algorithm
/***********

cases
- sequential repeat
  - checking the count of characters, whether they're equal
- change in value
  - save the 2 character pair
  - radially, look outwards for more valid sequences

sub strings == the min of the 2 lengths
update the total

prev = 0
// persist previous sequence length
current = 0
// this contains the current sequence value
currentSequence = '0'
for i < length(str) {
	// compare str[i] to the current sequence value
	// if repeat, update current sequence count

	// else, handle transition
	// total += min(prev, current)
	// prev = current
	// current = 0
	// set currentSequence to new character
}

************/

func countBinarySubstrings(s string) int {
	prev := 0
	current := 0
	total := 0
	currentSequence := byte('0')

	for i := 0; i <= len(s); i++ {
		if i < len(s) && s[i] == currentSequence {
			current++
		} else {
			if prev < current {
				total += prev
			} else {
				total += current
			}
			prev = current
			current = 1

			if i < len(s) {
				currentSequence = s[i]
			}
		}
	}

	return total
}
