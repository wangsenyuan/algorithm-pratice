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
	s := `2
0 0
0 1
0 2
1 0
1 1
1 2
2 0
2 1
2 2
`
	expect := []int{1, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
0 0
0 1
0 2
0 3
1 0
1 2
2 0
2 1
2 2
`
	expect := []int{0, 3}
	runSample(t, s, expect)
}
