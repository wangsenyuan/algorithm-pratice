package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	ask := func(l int, r int) int {
		if l == r {
			t.Fatalf("Sample ask wrong %d %d", l, r)
		}
		l--
		r--
		var c0 int
		var res int
		for i := l; i <= r; i++ {
			if s[i] == '1' {
				res += c0
			} else {
				c0++
			}
		}
		return res
	}

	res := solve(len(s), ask)

	if expect && res != s || !expect && res != "IMPOSSIBLE" {
		t.Fatalf("Sample expect %s, but got %s", s, res)
	}
}

func TestSample1(t *testing.T) {
	s := "01001"
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := "0000"
	runSample(t, s, false)
}

func TestSample3(t *testing.T) {
	s := "0100100010101111"
	runSample(t, s, true)
}

func TestSample4(t *testing.T) {
	s := "01"
	runSample(t, s, true)
}

func TestSample5(t *testing.T) {
	s := "010"
	runSample(t, s, true)
}

func TestSample6(t *testing.T) {
	s := "101"
	runSample(t, s, true)
}
