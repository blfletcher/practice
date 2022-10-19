package main

func minMovesToMakeBinaryPalindrome(s string) int {
	n := len(s)

	mismatches := 0
	zeroes, ones := 0, 0
	left, right := 0, n-1
	for left < right {
		if s[left] != s[right] {
			mismatches++
		}
		if s[left] == '0' {
			zeroes++
		} else {
			ones++
		}
		if s[right] == '0' {
			zeroes++
		} else {
			ones++
		}
		left++
		right--
	}

	if left == right {
		if s[left] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes%2 != 0 && ones%2 != 0 {
		return -1

	}

	if n%2 == 0 && (zeroes%2 != 0 || ones%2 != 0) {
		return -1
	}

	return (mismatches + 1) / 2
}
