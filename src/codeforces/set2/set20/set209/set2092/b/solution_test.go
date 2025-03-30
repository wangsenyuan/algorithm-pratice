package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("expect %v, but got %v", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `3
000
000`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
010001
010111`
	expect := true
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `5
10000
01010`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2
11
00`
	expect := true
	runSample(t, s, expect)
}