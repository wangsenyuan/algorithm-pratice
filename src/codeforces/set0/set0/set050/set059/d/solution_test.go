package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
5 4 1 2 6 3 7 8 9
5 6 2
9 3 4
1 7 8
4
`
	expect := []int{2, 3, 5, 6, 9, 1, 7, 8}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
5 4 1 2 6 3 7 8 9
5 6 2
9 3 4
1 7 8
8
`
	expect := []int{1, 2, 3, 4, 5, 6, 7, 9}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
4 1 3 2 5 6
4 6 5
1 2 3
4
`
	expect := []int{5, 6, 1, 2, 3}
	runSample(t, s, expect)
}
