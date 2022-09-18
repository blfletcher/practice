package main

import (
	"fmt"
	"testing"
)

var testCasesA = []testCaseIntStr{
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
	if true {
		t.Skip()
	}

	for i, v := range testCasesA {
		idx := i
		tst := v

		res := intToRoman(tst.input)
		fmt.Printf("\n\n[%d] [%d] [%s]\n\n\n", idx, tst.input, res)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%s]\t\tresult=[%s]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
