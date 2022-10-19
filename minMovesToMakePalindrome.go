package main

// https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/

func minMovesToMakePalindrome(s string) int {
	n := len(s)
	left, right := 0, n-1

	moves := 0
	center := -1

	modS := []rune(s)
	for left < right {
		if modS[left] == modS[right] {
			left++
			right--
			continue
		}

		j := left + 1
		for j < right {
			if modS[j] == modS[right] {
				break
			}
			j++
		}

		if j == right {
			center = right
			right--
			continue
		}

		for k := j; k > left; k-- {
			c := modS[k-1]
			modS[k-1] = modS[k]
			modS[k] = c
			moves++
		}

		left++
		right--
	}

	if center != -1 {
		moves += center - n/2
	}

	return moves
}
