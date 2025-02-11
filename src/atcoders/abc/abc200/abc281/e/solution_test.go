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
	s := `6 4 3
3 1 4 1 5 9
`
	expect := []int{5, 6, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 6 3
12 2 17 11 19 8 4 3 6 20
`
	expect := []int{21, 14, 15, 13, 13}
	runSample(t, s, expect)
}
