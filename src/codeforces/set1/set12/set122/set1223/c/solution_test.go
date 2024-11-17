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
		t.Fatalf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
100
50 1
49 1
100`
	expect := -1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
100 200 100 200 100 200 100 100
10 2
15 3
107`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1000000000 1000000000 1000000000
50 1
50 1
3000000000`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
200 100 100 100 100
69 5
31 2
90`
	expect := 4
	runSample(t, s, expect)
}
