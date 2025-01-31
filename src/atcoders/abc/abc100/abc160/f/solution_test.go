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
	s := `3
1 2
1 3
`
	expect := []int{2, 1, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 2
`
	expect := []int{1, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2
2 3
3 4
3 5
`
	expect := []int{2, 8, 12, 3, 3}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8
1 2
2 3
3 4
3 5
3 6
6 7
6 8
`
	expect := []int{40, 280, 840, 120, 120, 504, 72, 72}
	runSample(t, s, expect)
}
