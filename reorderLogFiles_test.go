package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseS struct {
	input    []string
	expected []string
}

var testCasesS = []testCaseS{
	{
		input:    []string{"1 n u", "r 527", "j 893", "6 14", "6 82"},
		expected: []string{"1 n u", "r 527", "j 893", "6 14", "6 82"},
	},
	{
		input:    []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
		expected: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
	},
	{
		input:    []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
		expected: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
	},
	{
		input:    []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act car"},
		expected: []string{"a8 act car", "g1 act car", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
	},
	{
		input:    []string{"dig1 8 1 5 1", "let1 art zero can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
		expected: []string{"let3 art zero", "let1 art zero can", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
	},
	{
		input:    []string{"wpylev6cnqv8 447241070789889321113517804297550370", "2syod 60098540876848105552318 69698830167476212 2", "iuw2x1r qmxealfvosqgkv yunonsq nxcuwudndrpsroty h"},
		expected: []string{"iuw2x1r qmxealfvosqgkv yunonsq nxcuwudndrpsroty h", "wpylev6cnqv8 447241070789889321113517804297550370", "2syod 60098540876848105552318 69698830167476212 2"},
	},
}

func Test_reorderLogFiles(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesS {
		idx := i
		tst := v

		res := reorderLogFiles(tst.input)
		if !reflect.DeepEqual(tst.expected, res) {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\nexpected\n%+v\n\nresult\n%+v\n\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
