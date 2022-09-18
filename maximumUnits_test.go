package main

import (
	"fmt"
	"testing"
)

type testCaseG struct {
	input    [][]int
	size     int
	expected int
}

var testCasesG = []testCaseG{
	{
		input: [][]int{
			{1, 3},
			{2, 2},
			{3, 1},
		},
		size:     4,
		expected: 8,
	},
	{
		input: [][]int{
			{5, 10},
			{2, 5},
			{4, 7},
			{3, 9},
		},
		size:     10,
		expected: 91,
	},
}

func Test_maximumUnits(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesG {
		idx := i
		tst := v

		res := maximumUnits(tst.input, tst.size)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
