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
	s := `4 2
BWWW
WBBW
WBBW
WWWB
`
	expect := 4
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `3 1
BWB
WWB
BWB
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 3
BWBBB
BWBBB
BBBBB
BBBBB
WBBBW
`
	expect := 2
	runSample(t, s, expect)
}
