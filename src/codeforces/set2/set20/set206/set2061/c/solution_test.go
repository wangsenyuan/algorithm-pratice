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
	s := `3
0 1 2`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
0 0 0 0 0`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
0 0 1 1 2`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
0 1 2 3 4`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
0 0 1 1 1`
	expect := 4
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `5
5 1 5 2 5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `1
0`
	expect := 2
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `4
2 3 1 1`
	expect := 0
	runSample(t, s, expect)
}
