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
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 5 10
2 15 5`
	expect := 25
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
2 93 78
1 71 59
3 57 96`
	expect := 221
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `19
19 23 16
5 90 13
12 85 70
19 67 78
12 16 60
18 48 28
5 4 24
12 97 97
4 57 87
19 91 74
18 100 76
7 86 46
9 100 57
3 76 73
6 84 93
1 6 84
11 75 94
19 15 3
12 11 34`
	expect := 1354
	runSample(t, s, expect)
}
