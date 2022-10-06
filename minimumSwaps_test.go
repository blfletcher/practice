package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesAD = []testCaseArrIntInt{
	{
		input:    []int{3, 4, 5, 5, 3, 1},
		expected: 6,
	},
	{
		input:    []int{9},
		expected: 0,
	},
}

func Test_minimumSwaps(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAD {
		idx := i
		tst := v

		res := minimumSwaps(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
