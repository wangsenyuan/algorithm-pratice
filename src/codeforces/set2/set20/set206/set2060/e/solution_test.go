package main

import (
	"bufio"
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(bytes.NewReader([]byte(s)))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2 1
1 2
2 3
1 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 1
1 2
1 2`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2 0
3 2
1 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 0 0`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 3 1
1 2
1 3
2 3
1 2`
	expect := 2
	runSample(t, s, expect)
}
