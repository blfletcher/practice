package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseI struct {
	input1   []int
	input2   int
	expected [][]int
}

var testCasesI = []testCaseI{
	{
		input1:   []int{2, 3, 6, 7},
		input2:   7,
		expected: [][]int{{2, 2, 3}, {7}},
	},
	{
		input1:   []int{2, 3, 5},
		input2:   8,
		expected: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
	},
	{
		input1:   []int{2},
		input2:   1,
		expected: [][]int{},
	},
}

func Test_combinationSum(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesI {
		idx := i
		tst := v

		res := combinationSum(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n[%d]\n\nexpected\n\t\t%+v\nresult\n\t\t%+v\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
