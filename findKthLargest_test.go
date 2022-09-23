package main

import (
	"fmt"
	"testing"
)

type testCaseP struct {
	input1   []int
	input2   int
	expected int
}

var testCasesP = []testCaseP{
	{
		input1:   []int{-1, -1},
		input2:   2,
		expected: -1,
	},
	{
		input1:   []int{2, 1},
		input2:   2,
		expected: 1,
	},
	{
		input1:   []int{3, 2, 1, 5, 6, 4},
		input2:   2,
		expected: 5,
	},
	{
		input1:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		input2:   4,
		expected: 4,
	},
}

func Test_findKthLargest(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesP {
		idx := i
		tst := v

		res := findKthLargest(tst.input1, tst.input2)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%d] %+v\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input2, tst.input1, tst.expected, res))
		}
	}
}
