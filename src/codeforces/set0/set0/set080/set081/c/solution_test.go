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
	s := `5
3 2
4 4 5 4 4
`
	expect := []int{1, 1, 2, 1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
2 2
3 5 4 5
`
	expect := []int{1, 1, 2, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 5
4 4 4 5 4 4
`
	expect := []int{2, 2, 2, 1, 2, 2}
	runSample(t, s, expect)
}
