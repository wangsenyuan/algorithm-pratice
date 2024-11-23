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
	s := `6 3
1 3 9 8 24 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `100 3
94 94 83 27 80 73 61 38 34 95 72 96 59 36 78 15 83 78 39 22 21 57 54 59 9 32 81 64 94 90 67 41 18 57 93 76 44 62 77 61 31 70 39 73 81 57 43 31 27 85 36 26 44 26 75 23 66 53 3 14 40 67 53 19 70 81 98 12 91 15 92 90 89 86 58 30 67 73 72 69 68 47 30 7 89 35 17 93 45 6 4 23 73 36 10 34 73 74 45 27`
	expect := 27
	runSample(t, s, expect)
}
