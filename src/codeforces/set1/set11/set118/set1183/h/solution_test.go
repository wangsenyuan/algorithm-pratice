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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
asdf
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 6
aaaaa
`
	// a, aa, aaa, aaaa, aaaaa
	expect := 15
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 7
aaaaa
`
	// a, aa, aaa, aaaa, aaaaa
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 100
ajihiushda
`
	// a, aa, aaa, aaaa, aaaaa
	expect := 233
	runSample(t, s, expect)
}
func TestSample5(t *testing.T) {
	s := `4 4
abba
`
	// abba
	// abb aba bba
	// ab aa bb ba
	// a b
	expect := 3
	runSample(t, s, expect)
}
