package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseV struct {
	input1   []int
	input2   []int
	expected int
}

var testCasesV = []testCaseV{
	{
		input1:   []int{1, 2, 3},
		input2:   []int{3, 2, 1},
		expected: 14,
	},
	{
		input1:   []int{-5, -3, -3, -2, 7, 1},
		input2:   []int{-10, -5, 3, 4, 6},
		expected: 102,
	},
}

func Test_maximumScore(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesV {
		idx := i
		tst := v

		res := maximumScore(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
