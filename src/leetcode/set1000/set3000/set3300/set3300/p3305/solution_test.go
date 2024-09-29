package p3305

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := countOfSubstrings(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aeioqq"
	k := 1
	expect := 0
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "aeiou"
	k := 0
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "ieaouqqieaouqq"
	k := 1
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "iqeaouqi"
	k := 2
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "eioaua"
	k := 0
	expect := 2
	runSample(t, s, k, expect)
}
