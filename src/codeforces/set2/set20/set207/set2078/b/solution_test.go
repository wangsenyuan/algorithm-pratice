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
+ 4 x 2
x 3 x 3
+ 7 + 4`
	expect := 32
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
+ 9 x 2
x 2 x 3
+ 9 + 10
x 2 + 1`
	expect := 98
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
x 2 + 1
+ 9 + 10
x 2 x 3
+ 9 x 2`
	expect := 144
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
x 3 x 3
x 2 x 2
+ 21 + 2
x 2 x 3
+ 41 x 3`
	expect := 351
	runSample(t, s, expect)
}
