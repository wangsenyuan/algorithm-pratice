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
	s := `3
110`
	expect := 500000006
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
100`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
1101001011`
	expect := 193359386
	runSample(t, s, expect)
}
