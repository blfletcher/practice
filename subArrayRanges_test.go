package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseAG struct {
	input    []int
	expected int64
}

var testCasesAG = []testCaseAG{
	{
		input:    []int{1, 2, 3},
		expected: 4,
	},
	{
		input:    []int{1, 3, 3},
		expected: 4,
	},
	{
		input:    []int{4, -2, -3, 4, 1},
		expected: 59,
	},
}

func Test_subArrayRanges(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAG {
		idx := i
		tst := v

		res := subArrayRanges(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
