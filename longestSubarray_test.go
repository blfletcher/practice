package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseAF struct {
	input1   []int
	input2   int
	expected int
}

var testCasesAF = []testCaseAF{
	{
		input1:   []int{24, 12, 71, 33, 5, 87, 10, 11, 3, 58, 2, 97, 97, 36, 32, 35, 15, 80, 24, 45, 38, 9, 22, 21, 33, 68, 22, 85, 35, 83, 92, 38, 59, 90, 42, 64, 61, 15, 4, 40, 50, 44, 54, 25, 34, 14, 33, 94, 66, 27, 78, 56, 3, 29, 3, 51, 19, 5, 93, 21, 58, 91, 65, 87, 55, 70, 29, 81, 89, 67, 58, 29, 68, 84, 4, 51, 87, 74, 42, 85, 81, 55, 8, 95, 39},
		input2:   87,
		expected: 25,
	},
	{
		input1:   []int{4, 10, 2, 6, 1},
		input2:   10,
		expected: 5,
	},
	{
		input1:   []int{8, 2, 4, 7},
		input2:   4,
		expected: 2,
	},
	{
		input1:   []int{10, 1, 2, 4, 7, 2},
		input2:   5,
		expected: 4,
	},
	{
		input1:   []int{4, 2, 2, 2, 4, 4, 2, 2},
		input2:   0,
		expected: 3,
	},
}

func Test_longestSubarray(t *testing.T) {
	if false {
		t.Skip()
	}

	for i, v := range testCasesAF {
		idx := i
		tst := v

		res := longestSubarray(tst.input1, tst.input2)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n[%d] %+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input2, tst.input1, tst.expected, res))
		}
	}
}
