package main

// https://leetcode.com/problems/integer-to-roman/

type numberPairs struct {
	oneChar  string
	fourChar string
	fiveChar string
	nineChar string
}

func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return "Invalid int"
	}

	intValues := map[int]numberPairs{
		1000: {
			oneChar: "M",
		},
		100: {
			oneChar:  "C",
			fourChar: "CD",
			fiveChar: "D",
			nineChar: "CM",
		},
		10: {
			oneChar:  "X",
			fourChar: "XL",
			fiveChar: "L",
			nineChar: "XC",
		},
		1: {
			oneChar:  "I",
			fourChar: "IV",
			fiveChar: "V",
			nineChar: "IX",
		},
	}

	currentVal := num
	romanVersion := ""
	for modifier := 1000; modifier > 0 && currentVal > 0; modifier /= 10 {
		charValue := currentVal / modifier
		currentVal = currentVal % modifier

		if charValue == 9 {
			romanVersion += intValues[modifier].nineChar
			charValue -= 9
		} else if charValue >= 5 {
			romanVersion += intValues[modifier].fiveChar
			charValue -= 5
		} else if charValue == 4 {
			romanVersion += intValues[modifier].fourChar
			charValue -= 4
		}

		for i := 0; i < charValue; i++ {
			romanVersion += intValues[modifier].oneChar
		}
	}

	return romanVersion
}
