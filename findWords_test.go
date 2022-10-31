package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type testCaseAL struct {
	input1   [][]byte
	input2   []string
	expected []string
}

var testCasesAL = []testCaseAL{
	{
		input1:   [][]byte{{'a', 'b', 'c', 'e'}, {'x', 'x', 'c', 'd'}, {'x', 'x', 'b', 'a'}},
		input2:   []string{"abc", "abcd"},
		expected: []string{"abc", "abcd"},
	},
	{
		input1:   [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
		input2:   []string{"oath", "pea", "eat", "rain"},
		expected: []string{"eat", "oath"},
	},
	{
		input1:   [][]byte{{'a', 'b'}, {'c', 'd'}},
		input2:   []string{"abcb"},
		expected: []string{},
	},
}

func Test_findWords(t *testing.T) {
	if false {
		t.Skip()
	}

	for i, v := range testCasesAL {
		idx := i
		tst := v

		res := findWords(tst.input1, tst.input2)
		sort.Slice(res, func(i, j int) bool {
			return res[i] > res[j]
		})
		sort.Slice(tst.expected, func(i, j int) bool {
			return tst.expected[i] > tst.expected[j]
		})
		if !reflect.DeepEqual(tst.expected, res) {
			msg := fmt.Sprintf("\n\ntest #%d\n", idx)
			msg += fmt.Sprintf("%+v\n%+v", tst.input1, tst.input2)
			msg += "\n\nexpected\n"
			msg += fmt.Sprintf("%+v", tst.expected)
			msg += "\n\nresult\n"
			msg += fmt.Sprintf("%+v", res)
			msg += "\n\n\n"
			t.Fatal(msg)
		}
	}
}
