package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseAB struct {
	input    []string
	expected [][]string
}

var testCasesAB = []testCaseAB{
	{
		input: []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
		expected: [][]string{
			{"acef"},
			{"a", "z"},
			{"abc", "bcd", "xyz"},
			{"az", "ba"},
		},
	},
	{
		input:    []string{"a"},
		expected: [][]string{{"a"}},
	},
}

func Test_groupStrings(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAB {
		idx := i
		tst := v

		res := groupStrings(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
