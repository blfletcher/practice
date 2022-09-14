package main

import (
	"fmt"
	"testing"
)

type testCase3 struct {
	input    int
	expected string
}

var testCases3 = []testCase3{
	{
		input:    3,
		expected: "III",
	},
	{
		input:    58,
		expected: "LVIII",
	},
	{
		input:    1994,
		expected: "MCMXCIV",
	},
	{
		input:    401,
		expected: "CDI",
	},
	{
		input:    899,
		expected: "DCCCXCIX",
	},
}

func Test_intToRoman(t *testing.T) {
	for i, v := range testCases3 {
		idx := i
		tst := v

		res := intToRoman(tst.input)
		fmt.Printf("\n\n[%d] [%d] [%s]\n\n\n", idx, tst.input, res)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%s]\t\tresult=[%s]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
