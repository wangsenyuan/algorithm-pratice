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
	s := `7 5
1 2 1
3 2 3
2 4 1
4 5 2
5 7 4
3 6 2
5 2 3 4 1
`
	expect := []int{21, 7, 15, 21, 3}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2
1 2
`
	expect := []int{0, 0}

	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 2 1
2 3 2
1 3 2
`
	expect := []int{1, 3, 3}

	runSample(t, s, expect)
}
