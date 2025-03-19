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
	s := `5
4 3 2 1 0`
	expect := 5
	runSample(t, s, expect)
}
func TestSample2(t *testing.T) {
	s := `6
4 3 3 2 1 0`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
2 0 1 2`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1
777`
	expect := 1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
1000000000 1 7 9`
	expect := 4
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `4
0 1 0 1`
	expect := 3
	runSample(t, s, expect)
}
