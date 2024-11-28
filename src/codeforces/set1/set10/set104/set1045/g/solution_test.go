package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
3 6 1
7 3 10
10 5 8
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 10
41 100000 1
78 100000 1
69 100000 4
58 100000 2
34 100000 3
0 100000 2
24 100000 2
64 100000 1
62 100000 1
67 100000 3
`
	expect := 45
	runSample(t, s, expect)
}
