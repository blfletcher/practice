package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    string
	expected int
}

var testCases = []testCase{
	{},
}

func Test_f(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCases {
		idx := i
		tst := v

		var res int
		f()
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d [%s]\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
