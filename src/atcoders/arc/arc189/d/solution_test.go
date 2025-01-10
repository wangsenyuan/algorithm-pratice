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
	s := `6
4 13 2 3 2 6
`
	expect := []int{4, 30, 2, 13, 2, 13}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `12
22 25 61 10 21 37 2 14 5 8 6 24
`
	expect := []int{22, 47, 235, 10, 31, 235, 2, 235, 5, 235, 6, 235}
	runSample(t, s, expect)
}
