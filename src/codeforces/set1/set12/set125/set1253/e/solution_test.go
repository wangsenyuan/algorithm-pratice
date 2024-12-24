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
	s := `3 595
43 2
300 4
554 10`
	expect := 281
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1
1 1
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 50
20 0
3 1
`
	expect := 30
	runSample(t, s, expect)
}
func TestSample4(t *testing.T) {
	s := `5 240
13 0
50 25
60 5
155 70
165 70
`
	expect := 26
	runSample(t, s, expect)
}
