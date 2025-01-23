package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := process(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	if expect == 0 {
		return
	}
	// expect == 3
	var b []int
	for _, i := range res {
		b = append(b, a[i-1])
	}
	for i := 1; i+1 < len(b); i++ {
		if b[i] > b[i-1] && b[i] > b[i+1] {
			return
		}
		if b[i] < b[i-1] && b[i] < b[i+1] {
			return
		}
	}
	t.Fatalf("Sample result %v, is ordered", b)
}

func TestSample1(t *testing.T) {
	s := `5
67 499 600 42 23
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 3 2
`
	expect := 3
	runSample(t, s, expect)
}