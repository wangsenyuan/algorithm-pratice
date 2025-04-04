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
	s := `5 8
2 3 1 1 2
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 10
1 3 4 2 1 4 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 15
5 5 5 5 5
`
	expect := 1
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `5 16
5 5 5 5 5
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 26
5 5 5 5 5
`
	expect := -1
	runSample(t, s, expect)
}