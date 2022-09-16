package main

import (
	"fmt"
	"testing"
)

type testCase6 struct {
	input    []int
	expected int
}

var testCases6 = []testCase6{
	{
		input:    []int{3, 1, 2, 4},
		expected: 17,
	},
	{
		input:    []int{11, 81, 94, 43, 3},
		expected: 444,
	},
}

func Test_sumSubarrayMins(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCases6 {
		idx := i
		tst := v

		res := sumSubarrayMins(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
