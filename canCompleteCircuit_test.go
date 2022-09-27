package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseY struct {
	input1   []int
	input2   []int
	expected int
}

var testCasesY = []testCaseY{
	{
		input1:   []int{5, 1, 2, 3, 4},
		input2:   []int{4, 4, 1, 5, 1},
		expected: 4,
	},
	{
		input1:   []int{1, 2, 3, 4, 5},
		input2:   []int{3, 4, 5, 1, 2},
		expected: 3,
	},
	{
		input1:   []int{2, 3, 4},
		input2:   []int{3, 4, 3},
		expected: -1,
	},
}

func Test_canCompleteCircuit(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesY {
		idx := i
		tst := v

		res := canCompleteCircuit(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
