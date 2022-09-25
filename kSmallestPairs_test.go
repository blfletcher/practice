package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseR struct {
	input1   []int
	input2   []int
	input3   int
	expected [][]int
}

var testCasesR = []testCaseR{
	{
		input1:   []int{1, 7, 11},
		input2:   []int{2, 4, 6},
		input3:   3,
		expected: [][]int{{1, 2}, {1, 4}, {1, 6}},
	},
	{
		input1:   []int{1, 1, 2},
		input2:   []int{1, 2, 3},
		input3:   2,
		expected: [][]int{{1, 1}, {1, 1}},
	},
	{
		input1:   []int{1, 2},
		input2:   []int{3},
		input3:   3,
		expected: [][]int{{1, 3}, {2, 3}},
	},
}

func Test_kSmallestPairs(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesR {
		idx := i
		tst := v
		res := kSmallestPairs(tst.input1, tst.input2, tst.input3)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n%+v\n[%d]\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input1, tst.input2, tst.input3, tst.expected, res))
		}
	}
}
