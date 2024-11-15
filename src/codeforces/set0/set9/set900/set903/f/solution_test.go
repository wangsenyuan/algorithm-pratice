package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	r := bufio.NewReader(strings.NewReader(s))

	res := process(r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 10 8 20
***.
***.
***.
...*
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
2 1 8 2
.***...
.***..*
.***...
....*..
`
	expect := 3
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := `4
10 10 1 10
***.
*..*
*..*
.***
`
	expect := 2
	runSample(t, s, expect)
}
