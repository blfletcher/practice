package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseP3 struct {
	input    *TreeNode
	expected [][]int
}

var testCasesP3 = []testCaseP3{
	{
		input: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: -1,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
		expected: [][]int{{0}, {4, 2}, {1, 3, -1}, {8, 6, 1, 5}},
	},
	{
		input: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		expected: [][]int{{1}, {3, 2}, {4, 5}},
	},
	{
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		expected: [][]int{{3}, {20, 9}, {15, 7}},
	},
	{
		input: &TreeNode{
			Val: 1,
		},
		expected: [][]int{{1}},
	},
	{
		input:    nil,
		expected: [][]int{},
	},
}

func Test_zigzagLevelOrder(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesP3 {
		idx := i
		tst := v

		res := zigzagLevelOrder(tst.input)
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
