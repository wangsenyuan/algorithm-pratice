package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, name, p := process(reader)
	if res == expect {
		return
	}
	if expect == fail || res == fail {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	m := len(name)
	for i := range len(p) {
		if p[i] == '1' {
			if res[i:i+m] != name {
				t.Fatalf("Sample result %s, not correct at %d", res, i)
			}
		} else {
			if res[i:i+m] == name {
				t.Fatalf("Sample result %s, not correct at %d", res, i)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
aba
101
`
	expect := "ababa"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 2
a
10001
`
	expect := "abbba"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 2
abba
101
`
	expect := fail
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `12 5
abacaba
010001
`
	expect := "aabacabacaba"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `20 5
ece
010101001001000010
`
	expect := "aecececeeceeceaaecea"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `15 2
aabb
010000000100
`
	expect := "aabacabacaba"
	runSample(t, s, expect)
}