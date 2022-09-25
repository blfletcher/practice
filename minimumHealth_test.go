package main

import (
	"fmt"
	"testing"
)

type testCaseQ struct {
	input1   []int
	input2   int
	expected int64
}

var testCasesQ = []testCaseQ{
	{
		input1:   []int{2, 7, 4, 3},
		input2:   4,
		expected: 13,
	},
	{
		input1:   []int{2, 5, 3, 4},
		input2:   7,
		expected: 10,
	},
	{
		input1:   []int{3, 3, 3},
		input2:   0,
		expected: 10,
	},
}

func Test_minimumHealth(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesQ {
		idx := i
		tst := v

		// replace
		res := minimumHealth(tst.input1, tst.input2)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n[%d]\n\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
