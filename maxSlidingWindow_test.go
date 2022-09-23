package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseM struct {
	input1   []int
	input2   int
	expected []int
}

var testCasesM = []testCaseM{
	{
		input1:   []int{1, 3, -1, -3, 5, 3, 6, 7},
		input2:   3,
		expected: []int{3, 3, 5, 5, 6, 7},
	},
	{
		input1:   []int{1},
		input2:   1,
		expected: []int{1},
	},
	{
		input1:   []int{-1, -2, -3},
		input2:   3,
		expected: []int{-1},
	},
}

func Test_maxSlidingWindow(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesM {
		idx := i
		tst := v

		res := maxSlidingWindow(tst.input1, tst.input2)
		if !reflect.DeepEqual(res, tst.expected) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%d] %+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input2, tst.input1, tst.expected, res))
		}
	}
}
