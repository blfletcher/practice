package main

import (
	"fmt"
	"testing"
)

var testCasesC = []testCaseArrIntInt{
	{
		input:    []int{1, 3, 1, 2},
		expected: 44,
	},
	{
		input:    []int{5, 4, 6},
		expected: 213,
	},
	{
		input:    []int{747, 812, 112, 1230, 1426, 1477, 1388, 976, 849, 1431, 1885, 1845, 1070, 1980, 280, 1075, 232, 1330, 1868, 1696, 1361, 1822, 524, 1899, 1904, 538, 731, 985, 279, 1608, 1558, 930, 1232, 1497, 875, 1850, 1173, 805, 1720, 33, 233, 330, 1429, 1688, 281, 362, 1963, 927, 1688, 256, 1594, 1823, 743, 553, 1633, 1898, 1101, 1278, 717, 522, 1926, 1451, 119, 1283, 1016, 194, 780, 1436, 1233, 710, 1608, 523, 874, 1779, 1822, 134, 1984},
		expected: 471441678,
	},
}

func Test_totalStrength(t *testing.T) {
	if true {
		t.Skip()
	}

	for i, v := range testCasesC {
		idx := i
		tst := v

		res := totalStrength(tst.input)
		if tst.expected != res {
			t.Fatal(fmt.Errorf("\n\ntest #%d\n%+v\n\t\texpected=[%d]\t\tresult=[%d]\n\n", idx, tst.input, tst.expected, res))
		}
	}
}
