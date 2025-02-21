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
	s := `4
1 3 4 2
`
	expect := 20
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
5 5 6 1 1 1
`
	// 6 1 1 1 5 5
	// 貌似是可以达到 58的
	// 6 * 2 + 5 * 4 + 5 * 4 + 1 * 3 + 2 + 1
	// 12 + 40 + 6 = 58
	expect := 58
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
8 6 9 1 2 1
`
	expect := 85
	runSample(t, s, expect)
}
