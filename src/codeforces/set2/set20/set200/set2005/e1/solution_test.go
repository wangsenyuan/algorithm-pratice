package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2 3
1 2
1 3 5
4 5 2`
	expect := "N"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 4
1 2
1 1 3 2
4 2 5 1`
	expect := "T"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 4 2
1 2
3 4
5 5
5 5
5 5`
	expect := "N"
	runSample(t, s, expect)
}
