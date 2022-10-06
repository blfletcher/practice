package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseAC struct {
	input1   []int
	input2   int
	expected int
}

var testCasesAC = []testCaseAC{
	{
		input1:   []int{1, 2, 5},
		input2:   11,
		expected: 3,
	},
	{
		input1:   []int{2},
		input2:   3,
		expected: -1,
	},
	{
		input1:   []int{1},
		input2:   0,
		expected: 0,
	},
}

func Test_coinChange(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAC {
		idx := i
		tst := v

		res := coinChange(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%d] %+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input2, tst.input1, tst.expected, res))
		}
	}
}
