package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesX = []testCaseStrInt{
	{
		input:    "aacaba",
		expected: 2,
	},
	{
		input:    "abcd",
		expected: 1,
	},
}

func Test_numSplits(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesX {
		idx := i
		tst := v

		res := numSplits(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%s]\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
