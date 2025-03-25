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
	s := `8 25
4 5 18 3 17 17 18 14`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 21
20 14 1 4 20 8 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 1
20 5 9 4 14 12 2 20`
	expect := 7
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 37
2 13 13 11 12 19 16 18`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 38
15 3 14 7`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `2 3
1 4`
	expect := 0
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `1 1
1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `3 10
2 7 11`
	expect := 0
	runSample(t, s, expect)
}
