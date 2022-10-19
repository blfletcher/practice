package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseP1 struct {
	input    []int
	expected int
}

var testCasesP1 = []testCaseP1{
	{
		input:    []int{},
		expected: 1,
	},
	{
		input:    []int{},
		expected: 0,
	},
	{
		input:    []int{},
		expected: 0,
	},
}

func Test_p1(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesP1 {
		idx := i
		tst := v

		res := p1(tst.input)
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
