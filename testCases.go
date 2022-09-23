package main

type testCaseIntBool struct {
	input    int
	expected bool
}

type testCaseIntStr struct {
	input    int
	expected string
}

type testCaseStrInt struct {
	input    string
	expected int
}

type testCaseDArrIntInt struct {
	input    [][]int
	expected int
}

type testCaseDArrByteInt struct {
	input    [][]byte
	expected int
}

type testCaseArrIntInt struct {
	input    []int
	expected int
}

type testCaseArrStrInt struct {
	input    []string
	expected int
}
