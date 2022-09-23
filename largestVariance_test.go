package main

import (
	"fmt"
	"testing"
)

var testCasesL = []testCaseStrInt{
	{
		input:    "aababbb",
		expected: 3,
	},
	{
		input:    "abcde",
		expected: 0,
	},
}

func Test_largestVariance(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesL {
		idx := i
		tst := v

		res := largestVariance(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%s]\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
