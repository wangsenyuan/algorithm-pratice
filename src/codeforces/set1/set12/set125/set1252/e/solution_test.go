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
	s := `16 1 8 6
1 3 5 6 7 8 9 10 3 7 8 9 10 11 12 12
`
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 2, 3, 4, 5, 6, 7, 8, 8}

	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `16 1 8 6
1 3 5 6 7 8 9 10 3 7 8 9 10 11 12 13
`
	runSample(t, s, nil)
}

func TestSample3(t *testing.T) {
	s := `16 1 10 10
1 3 5 6 7 8 9 10 3 7 8 9 1 11 12 13
`
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 1, 2, 3, 4}

	runSample(t, s, expect)
}
