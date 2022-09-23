package main

import (
	"fmt"
	"testing"
)

type testCaseQ struct {
	input    []int
	expected int
}

var testCasesQ = []testCaseQ{
	{
		input:    []int{},
		expected: 0,
	},
	{
		input:    []int{},
		expected: 0,
	},
}

func Test_f1(t *testing.T) {
	// remove
	if true {
		t.Skip()
	}

	for i, v := range testCasesQ {
		idx := i
		tst := v

		// replace
		res := f1(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
