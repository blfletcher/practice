package main

import (
	"fmt"
	"testing"
)

var testCasesK = []testCaseDArrByteInt{
	{
		input: [][]byte{
			{'X', 'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'O', 'O', 'O', 'X'},
			{'X', 'O', 'O', '#', 'O', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'X'},
		},
		expected: 3,
	},
	{
		input: [][]byte{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'X', 'O', 'X'},
			{'X', 'O', 'X', '#', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
		},
		expected: -1,
	},
	{
		input: [][]byte{
			{'O', '*'},
		},
		expected: -1,
	},
}

func Test_getFood(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesK {
		idx := i
		tst := v

		res := getFood(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%s]\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
