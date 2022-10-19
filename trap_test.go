package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesAJ = []testCaseArrIntInt{
	{
		input:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		expected: 6,
	},
	{
		input:    []int{4, 2, 0, 3, 2, 5},
		expected: 9,
	},
}

func Test_trap(t *testing.T) {
	if false {
		t.Skip()
	}

	for i, v := range testCasesAJ {
		idx := i
		tst := v

		res := trap(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			msg := fmt.Sprintf("\n\ntest #%d\n", idx)
			msg += fmt.Sprintf("%+v", tst.input)
			msg += "\n\nexpected\n"
			msg += fmt.Sprintf("%+v", tst.expected)
			msg += "\n\nresult\n"
			msg += fmt.Sprintf("%+v", res)
			msg += "\n\n\n"
			t.Fatal(msg)
		}
	}
}
