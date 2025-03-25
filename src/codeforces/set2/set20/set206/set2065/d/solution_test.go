package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
4 4
6 1`
	expect := 41
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := `3 4
2 2 2 2
3 2 1 2
4 1 2 1`
	expect := 162
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 3
3 4 5
1 1 9`
	expect := 72
	runSample(t, s, expect)
}
