package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseW struct {
	input1   string
	input2   [][]int
	expected []int
}

var testCasesW = []testCaseW{
	{
		input1:   "**|**|***|",
		input2:   [][]int{{2, 5}, {5, 9}},
		expected: []int{2, 3},
	},
	{
		input1:   "***|**|*****|**||**|*",
		input2:   [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
		expected: []int{9, 0, 0, 0, 0},
	},
}

func Test_platesBetweenCandles(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesW {
		idx := i
		tst := v

		res := platesBetweenCandles(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%s] %+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
