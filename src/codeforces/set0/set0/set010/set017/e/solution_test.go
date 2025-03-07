package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	res := solve1(s)
	expect := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "babb"
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := "ciicbeefifb"
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := strings.Repeat("a", 97453)
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := strings.Repeat("a", 1785334)
	runSample(t, s)
}
