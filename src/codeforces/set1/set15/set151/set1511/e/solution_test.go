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
	s := `3 4
**oo
oo*o
**oo
`
	expect := 144
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
**oo
oo**
**oo
`
	expect := 48
	runSample(t, s, expect)
}
