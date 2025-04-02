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
	s := `FT
1
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `FFFTFFF
2
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
50
`
	expect := 100
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `TFFTFTTTFFTFTFTTTFFTTFFTFTFFTFFFFTTTTTFTFTFTTFFTTFTFFTTFTFFTTTTFFTFTTTFTTTTFFFFTFFTFFTFFFFTTTT
2
`
	expect := 19
	runSample(t, s, expect)
}
