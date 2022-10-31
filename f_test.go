package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseN struct {
	input    []int
	expected int
}

var testCasesN = []testCaseN{
	{
		input:    []int{},
		expected: 0,
	},
	{
		input:    []int{},
		expected: 0,
	},
}

func Test_fN(t *testing.T) {
	// remove
	if true {
		t.Skip()
	}

	// replace
	for i, v := range testCasesN {
		idx := i
		tst := v

		// replace
		res := fN(tst.input)
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
