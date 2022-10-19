package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesAI = []testCaseStrInt{
	{
		input:    "0011",
		expected: 1,
	},
	{
		input:    "01001",
		expected: 1,
	},
	{
		input:    "0111",
		expected: -1,
	},
	{
		input:    "0100101001",
		expected: 2,
	},
	{
		input:    "010110011",
		expected: 1,
	},
	{
		input:    "010110101",
		expected: 2,
	},
	{
		input:    "010101011",
		expected: 1,
	},
	{
		input:    "010101101",
		expected: 2,
	},
	{
		input:    "111000101",
		expected: 1,
	},
}

func Test_minMovesToMakeBinaryPalindrome(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAI {
		idx := i
		tst := v

		res := minMovesToMakeBinaryPalindrome(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
