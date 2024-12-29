package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	words := strings.Split(s, " ")
	res := solve(words)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "I want to order pizza"
	expect := "Iwantorderpizza"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "sample please ease in out"
	expect := "sampleaseinout"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abc abc abc abc abc"
	expect := "abc"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "abc123 123abc abc abc abc"
	expect := "abc123abc"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "999 9999"
	expect := "9999"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "1101 1001 001001 101 010"
	expect := "1101001001010"
	runSample(t, s, expect)
}
