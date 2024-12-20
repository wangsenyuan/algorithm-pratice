package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
2 3 4`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
3 3`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
-1 -1 -1`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
3 4 -1`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
4 -1 4 5`
	expect := true
	runSample(t, s, expect)
}
