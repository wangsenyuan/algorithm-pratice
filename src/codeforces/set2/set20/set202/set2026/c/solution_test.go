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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
1`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
101101`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
1110001`
	expect := 18
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
11111`
	expect := 6
	runSample(t, s, expect)
}
