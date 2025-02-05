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
	s := `6
1 1 1
2 2 0
2 2 1
1 3 0
2 2 0
2 2 2
`
	expect := []int{0, 1, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 1 0
2 2 0
2 0 3
1 0 2
2 1 3
2 1 6
`
	expect := []int{2, 2, 3, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
1 1 2
1 2 3
2 3 3
1 0 0
1 5 1
2 5 0
2 4 0
`
	expect := []int{1, 1, 2}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7
1 1 3
1 2 3
2 3 4
1 2 0
1 5 3
2 5 5
2 7 22
`
	expect := []int{1, 2, 3}
	runSample(t, s, expect)
}
