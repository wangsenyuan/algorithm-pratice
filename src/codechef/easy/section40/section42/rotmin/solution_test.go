package main

import "testing"

func runSample(t *testing.T, s string, p int, q int, expect string) {
	res := solve(s, p, q)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p, q := 1, 2
	s := "bazooka"
	expect := "aaanoka"
	runSample(t, s, p, q, expect)
}

func TestSample2(t *testing.T) {
	p, q := 5, 2
	s := "sdlftcaasd"
	expect := "qdlftcaasd"
	runSample(t, s, p, q, expect)
}

func TestSample3(t *testing.T) {
	p, q := 2, 4
	s := "adivjnaefv"
	expect := "aahvjnaefv"
	runSample(t, s, p, q, expect)
}

func TestSample4(t *testing.T) {
	p, q := 7, 1
	s := "lakjsdcedg"
	expect := "kakjsdcedg"
	runSample(t, s, p, q, expect)
}
