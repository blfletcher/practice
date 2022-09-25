package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesU = []testCaseArrIntInt{
	{
		input:    []int{3, 4, 2},
		expected: 6,
	},
	{
		input:    []int{2, 2, 3, 3, 3, 4},
		expected: 9,
	},
}

func Test_deleteAndEarn(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesU {
		idx := i
		tst := v

		res := deleteAndEarn(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
