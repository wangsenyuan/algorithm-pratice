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
	s := `100 200 250 150 200 250`
	expect := "Ron"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `100 50 50 200 200 100`
	expect := "Hermione"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 10 200 20 300 30`
	expect := "Hermione"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `0 0 0 0 0 0`
	expect := "Hermione"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1 1 0 1 1 1`
	expect := "Ron"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `1 0 1 2 1 2`
	expect := "Hermione"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `100 1 100 1 0 1`
	expect := "Ron"
	runSample(t, s, expect)
}
func TestSample8(t *testing.T) {
	s := `0 0 0 493 0 0`
	expect := "Ron"
	runSample(t, s, expect)
}
