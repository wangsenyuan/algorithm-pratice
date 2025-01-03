package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2 10 4
0 1
1 1
0 4
1 2
`
	expect := [][]int{
		{4, 1},
		{5, 2},
		{4, 2},
		{5, 3},
	}
	runSample(t, s, expect)
}
