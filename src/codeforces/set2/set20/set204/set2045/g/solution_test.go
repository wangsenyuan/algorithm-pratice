package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 1
3359
4294
3681
5
1 1 3 4
3 3 2 1
2 2 1 4
1 3 3 2
1 1 1 1
`
	expect := []string{
		"2",
		"4",
		"-7",
		"-1",
		"0",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4 5
1908
2023
2
1 1 2 4
1 1 1 1
`
	expect := []string{
		"INVALID",
		"INVALID",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3 9
135
357
579
2
3 3 1 1
2 2 2 2
`
	expect := []string{
		"2048",
		"0",
	}
	runSample(t, s, expect)
}
