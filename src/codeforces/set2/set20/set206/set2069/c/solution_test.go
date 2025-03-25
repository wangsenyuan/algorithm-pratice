package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7
3 2 1 2 2 1 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9
1 2 3 2 1 3 2 2 3`
	expect := 22
	runSample(t, s, expect)
}
