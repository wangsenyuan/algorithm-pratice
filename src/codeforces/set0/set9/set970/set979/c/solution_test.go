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
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1 3
1 2
2 3`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 3
1 2
1 3`
	expect := 4
	runSample(t, s, expect)
}
