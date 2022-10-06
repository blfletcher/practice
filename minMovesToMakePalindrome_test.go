package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesAH = []testCaseStrInt{
	{
		input:    "aabb",
		expected: 2,
	},
	{
		input:    "letelt",
		expected: 2,
	},
	{
		input:    "cbaba",
		expected: 3,
	},
}

func Test_minMovesToMakePalindrome(t *testing.T) {
	if false {
		t.Skip()
	}

	for i, v := range testCasesAH {
		idx := i
		tst := v

		// replace
		res := minMovesToMakePalindrome(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
