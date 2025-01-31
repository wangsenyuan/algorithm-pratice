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
	s := "4 50 2 10 1 20 2 30 1"
	expect := 80
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "7 20 1 2 1 10 3 100 2 8 2 5 20 50 10"
	expect := 185
	runSample(t, s, expect)
}
