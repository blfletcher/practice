package main

import (
	"fmt"
	"testing"
)

var testCasesF = []testCaseArrIntInt{
	{
		input:    []int{1, 5, 0, 3, 5},
		expected: 3,
	},
	{
		input:    []int{},
		expected: 0,
	},
}

func Test_minimumOperations(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesF {
		idx := i
		tst := v

		res := minimumOperations(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
