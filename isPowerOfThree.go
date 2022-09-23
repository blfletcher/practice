package main

import (
	"math"
	"strconv"
)

// https://leetcode.com/problems/power-of-three/

func isPowerOfThree_v1(n int) bool {
	if n == 1 || n == 3 {
		return true
	}
	if n%3 != 0 || n < 3 {
		return false
	}

	i := 1
	pow := math.Pow(3, float64(i))
	n /= int(pow)
	for int(pow) < n {
		if n%3 != 0 {
			return false
		}

		pow = math.Pow(3, float64(i))
		i++
	}

	return n/int(pow) == 1 && n%int(pow) == 0
}

func isPowerOfThree(n int) bool {
	base3 := strconv.FormatInt(int64(n), 3)
	if len(base3) < 1 || base3[0] != '1' {
		return false
	}

	for i := 1; i < len(base3); i++ {
		if base3[i] != '0' {
			return false
		}
	}

	return true
}
