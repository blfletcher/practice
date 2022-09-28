package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseZ struct {
	input1   int
	input2   [][]int
	expected bool
}

var testCasesZ = []testCaseZ{
	{
		input1:   2,
		input2:   [][]int{{0, 1}},
		expected: true,
	},
	{
		input1:   2,
		input2:   [][]int{{1, 0}},
		expected: true,
	},
	{
		input1:   2,
		input2:   [][]int{{1, 0}, {0, 1}},
		expected: false,
	},
}

func Test_canFinish(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesZ {
		idx := i
		tst := v

		res := canFinish(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%d] %+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
