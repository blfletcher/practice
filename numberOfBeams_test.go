package main

import (
	"fmt"
	"testing"
)

type testCase7 struct {
	input    []string
	expected int
}

var testCases7 = []testCase7{
	{
		input:    []string{"011001", "000000", "010100", "001000"},
		expected: 8,
	},
	{
		input:    []string{"000", "111", "000"},
		expected: 0,
	},
}

func Test_numberOfBeams(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCases7 {
		idx := i
		tst := v

		res := numberOfBeams(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d [%s]\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
