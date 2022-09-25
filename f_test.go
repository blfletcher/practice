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
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}