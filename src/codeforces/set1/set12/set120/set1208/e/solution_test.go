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
	s := `3 3
3 2 4 8
2 2 5
2 6 3
`
	expect := []int{10, 15, 16}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
2 7 8
1 -8
`
	expect := []int{7, 8}
	runSample(t, s, expect)
}
