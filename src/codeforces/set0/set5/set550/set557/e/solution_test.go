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
	s := `abbabaab
7
`
	expect := "abaa"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `aaaaa
10
`
	expect := "aaa"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `bbaabb
13
`
	expect := "bbaabb"
	runSample(t, s, expect)
}
