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
1 7 2 3
8 1 1 3
`
	expect := []int{2, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
1 2 1 2 1 2 1 2
2 1 2 1 2 1 2 1
`
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
	runSample(t, s, expect)
}
