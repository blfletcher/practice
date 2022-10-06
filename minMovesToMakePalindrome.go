package main

// https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/

func minMovesToMakePalindrome(s string) int {
	// equal front to reverse
	// if odd length, can have 1 odd value at center
	// if even length, all characters must be even count
	// start from center and swap to make values equal on both sides
	// you can search for character index to swap

	// iterate 1 time, save character counts, positions
	//   map[char][]int

	// iterate character map
	//   sort the position array by distance from center
	//   if odd, note character as odd

	// moves = 0
	// if odd, left, right := center
	// else left,right:=floor(center),ceiling(center)
	// iterate the string while left >= 0 and right < len(string)
	//   if odd and left==right, move closest odd to center while incrementing moves
	//   else if left char != right char, find closest left vs right and move accordingly while incrementing moves
	//   left--
	//   right++

	charMap := map[rune][]int{}
	for i := range s {
		c := rune(s[i])

		if _, ok := charMap[c]; ok {
			charMap[c] = append(charMap[c], i)
		} else {
			charMap[c] = []int{i}
		}
	}

	n := len(s)
	left, right := 0, n-1

	moves := 0
	modS := []rune(s)
	for left < right {
		if modS[left] != modS[right] {
			leftPos := charMap[modS[right]]
			closestLeft := 0
			for leftPos[closestLeft] < left {
				closestLeft++
			}

			rightPos := charMap[modS[left]]
			closestRight := len(rightPos) - 1
			for rightPos[closestRight] > right {
				closestRight--
			}

			if (rightPos[closestRight] - right) > (leftPos[closestLeft] - left) {
				cc := modS[leftPos[closestLeft]]
				modS[leftPos[closestLeft]] = modS[left]
				modS[left] = cc
				newIdx := charMap[modS[right]][closestLeft]
				charMap[modS[right]][closestLeft] = left
				moves++
				for i := left + 1; i <= newIdx; i++ {
					currChar := modS[i]
					for j := range charMap[currChar] {
						if charMap[currChar][j] == i-1 {
							charMap[currChar][j] = i
							moves++
							break
						}
					}
				}
			} else {
				cc := modS[rightPos[closestRight]]
				modS[rightPos[closestRight]] = modS[right]
				modS[right] = cc
				newIdx := charMap[modS[left]][closestRight]
				charMap[modS[left]][closestRight] = right
				moves++
				for i := right - 1; i >= newIdx; i-- {
					currChar := modS[i]
					for j := len(charMap[currChar]) - 1; j >= 0; j-- {
						if charMap[currChar][j] == i-1 {
							charMap[currChar][j] = i
							moves++
							break
						}
					}
				}
			}
		}
		left++
		right--
	}

	return moves
}
