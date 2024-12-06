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
	s := `6
a := h
b := aha
c = a + b
c = c + c
e = c + c
d = a + c`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `15
x := haha
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x
x = x + x`
	expect := 32767
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
haha := hah`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
haahh := aaaha
ahhhh = haahh + haahh
haahh = haahh + haahh
ahhhh = ahhhh + haahh
ahhaa = haahh + ahhhh`
	expect := 0
	runSample(t, s, expect)
}
