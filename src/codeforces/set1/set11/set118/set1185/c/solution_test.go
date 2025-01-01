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
	s := `7 15
1 2 3 4 5 6 7
`
	expect := []int{0, 0, 0, 0, 0, 2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 100
80 40 40 40 60
`
	expect := []int{0, 1, 1, 2, 3}
	runSample(t, s, expect)
}
