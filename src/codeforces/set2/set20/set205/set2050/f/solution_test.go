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
	s := `5 5
5 14 2 6 3
4 5
1 4
2 4
3 5
1 1`
	expect := []int{3, 1, 4, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1
7
1 1`
	expect := []int{0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2
1 7 8
2 3
1 2`
	expect := []int{1, 6}
	runSample(t, s, expect)
}
