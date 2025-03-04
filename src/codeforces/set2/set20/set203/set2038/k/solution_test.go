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
	s := `4 2 4`
	expect := 21
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 210 420`
	expect := 125
	runSample(t, s, expect)
}
