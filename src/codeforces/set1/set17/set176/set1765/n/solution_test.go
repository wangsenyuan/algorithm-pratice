package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10000"
	k := 4
	expect := "1"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1337"
	k := 0
	expect := "1337"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "987654321"
	k := 6
	expect := "321"
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "66837494128"
	k := 5
	expect := "344128"
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "7808652"
	k := 3
	expect := "7052"
	runSample(t, s, k, expect)
}

func TestSample6(t *testing.T) {
	s := "706050"
	k := 4
	expect := "50"
	runSample(t, s, k, expect)
}
