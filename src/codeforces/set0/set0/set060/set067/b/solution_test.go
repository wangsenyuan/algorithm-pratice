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
	s := `5 2
1 2 1 0 0
`
	expect := []int{4, 1, 5, 2, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 0 0 0
`
	expect := []int{2, 3, 1, 4}
	runSample(t, s, expect)
}
