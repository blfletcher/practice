package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseT struct {
	input1   string
	input2   string
	expected int
}

var testCasesT = []testCaseT{
	{
		input1:   "ylqpejqbalahwr",
		input2:   "yrkzavgdmdgtqpg",
		expected: 3,
	},
	{
		input1:   "abcde",
		input2:   "ace",
		expected: 3,
	},
	{
		input1:   "abc",
		input2:   "abc",
		expected: 3,
	},
	{
		input1:   "abc",
		input2:   "def",
		expected: 0,
	},
	{
		input1:   "abcaba",
		input2:   "abba",
		expected: 4,
	},
	{
		input1:   "ababc",
		input2:   "abba",
		expected: 3,
	},
}

func Test_longestCommonSubsequence(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesT {
		idx := i
		tst := v
		res := longestCommonSubsequence(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%s] [%s]\n\nexpected\n[%d]\n\nresult\n[%d]\n\n\n", idx, tst.input1, tst.input2, tst.expected, res))
		}
	}
}
