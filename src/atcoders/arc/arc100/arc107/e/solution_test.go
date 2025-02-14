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
	s := `4
1 2 0 2
0
0
0
`
	expect := []int{7, 4, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `20
1 1 2 2 1 2 0 2 2 1 0 2 0 2 0 2 1 1 1 1
0
1
0
2
1
0
0
0
2
1
1
2
2
1
2
0
1
1
0
`
	expect := []int{176, 175, 49}
	runSample(t, s, expect)
}
