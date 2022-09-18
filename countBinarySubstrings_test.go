package main

import (
	"fmt"
	"testing"
)

var testCasesH = []testCaseStrInt{
	{
		input:    "00110011",
		expected: 6,
	},
	{
		input:    "10101",
		expected: 4,
	},
}

func Test_countBinarySubstrings(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesH {
		idx := i
		tst := v

		res := countBinarySubstrings(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d [%s]\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
