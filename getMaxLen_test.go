package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesAE = []testCaseArrIntInt{
	{
		input:    []int{1, -2, -3, 4},
		expected: 4,
	},
	{
		input:    []int{0, 1, -2, -3, -4},
		expected: 3,
	},
	{
		input:    []int{-1, -2, -3, 0, 1},
		expected: 2,
	},
}

func Test_getMaxLen(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAE {
		idx := i
		tst := v

		res := getMaxLen(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
