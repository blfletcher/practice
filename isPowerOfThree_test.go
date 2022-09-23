package main

import (
	"fmt"
	"testing"
)

var testCasesO = []testCaseIntBool{
	{
		input:    45,
		expected: false,
	},
	{
		input:    27,
		expected: true,
	},
	{
		input:    0,
		expected: false,
	},
	{
		input:    -1,
		expected: false,
	},
	{
		input:    6,
		expected: false,
	},
	{
		input:    59049,
		expected: true,
	},
}

func Test_isPowerOfThree(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesO {
		idx := i
		tst := v

		res := isPowerOfThree(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d [%d]\n\t\texpected=[%t]\t\tresult=[%t]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
