package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var testCasesAK = []testCaseStrArrStr{
	{
		input:    "23",
		expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
	{
		input:    "",
		expected: []string{},
	},
	{
		input:    "2",
		expected: []string{"a", "b", "c"},
	},
	{
		input:    "27",
		expected: []string{"ap", "aq", "ar", "as", "bp", "bq", "br", "bs", "cp", "cq", "cr", "cs"},
	},
}

func Test_letterCombinations(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAK {
		idx := i
		tst := v

		res := letterCombinations(tst.input)
		sort.Slice(res, func(i, j int) bool {
			return res[i] > res[j]
		})
		sort.Slice(tst.expected, func(i, j int) bool {
			return tst.expected[i] > tst.expected[j]
		})
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
