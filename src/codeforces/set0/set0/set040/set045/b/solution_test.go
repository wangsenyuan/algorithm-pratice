package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
2 3 1
1 2 3 4
1 2 3 4
`
	expect := []int{1, 1, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 6
7 6 4 2 3 5 5 7
10 4 3 8 9 1
1 1 1 2 2 2
`
	expect := []int{1, 1, 1, 2, 1, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 6
2 1 2 3 4 5 6 7
10 4 3 8 9 1
1 1 1 1 1 1
`
	expect := []int{1, 1, 1, 1, 0, 0}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `25 4
14 12 6 19 13 19 15 25 25 3 17 2 5 23 6 20 7 2 14 6 4 24 16 11 21
8811002 2094993 3283083 1356127
55 9 24 47
`
	expect := []int{2, 6, 0, 0}
	runSample(t, s, expect)
}
