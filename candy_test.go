package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCasesAA = []testCaseArrIntInt{
	{
		input:    []int{1, 0, 2},
		expected: 5,
	},
	{
		input:    []int{1, 2, 2},
		expected: 4,
	},
}

func Test_candy(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesAA {
		idx := i
		tst := v

		res := candy(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
