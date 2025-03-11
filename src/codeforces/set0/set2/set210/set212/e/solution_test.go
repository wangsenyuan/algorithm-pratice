package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []pair) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
2 3
3 4
4 5
`
	expect := []pair{
		{1, 3},
		{2, 2},
		{3, 1},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
1 2
2 3
3 4
5 6
6 7
7 4
8 9
9 10
10 4
`
	expect := []pair{
		{1, 8},
		{2, 7},
		{3, 6},
		{6, 3},
		{7, 2},
		{8, 1},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
3 1
2 1`
	expect := []pair{{1, 1}}
	runSample(t, s, expect)
}
