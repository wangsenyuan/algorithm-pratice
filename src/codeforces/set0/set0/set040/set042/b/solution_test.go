package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	chess := strings.Split(s, " ")
	res := solve(chess)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `a6 b4 c8 a8`
	expect := "CHECKMATE"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `a6 c4 b6 b8`
	expect := "OTHER"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `a2 b1 a3 a1`
	expect := "OTHER"
	runSample(t, s, expect)
}
func TestSample4(t *testing.T) {
	s := `a5 c5 c2 a1`
	expect := "CHECKMATE"
	runSample(t, s, expect)
}
