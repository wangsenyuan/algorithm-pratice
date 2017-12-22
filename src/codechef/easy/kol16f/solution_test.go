package main

import "testing"

func runSample(t *testing.T, chain []byte, expect int) {
	res := solve(chain)
	if res != expect {
		t.Errorf("sample: %s, should give %d, but got %d", chain, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []byte("RRRAAAARAA"), 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []byte("ARRRRAAA"), 4)
}

func TestSample3(t *testing.T) {
	runSample(t, []byte("AAAAAAAA"), 7)
}

func TestSample4(t *testing.T) {
	runSample(t, []byte("AB"), 2)
}

func TestSample5(t *testing.T) {
	runSample(t, []byte("ABA"), 2)
}

func TestSample6(t *testing.T) {
	runSample(t, []byte("A"), 1)
}

func TestSample7(t *testing.T) {
	runSample(t, []byte("AB"), 2)
}

func TestSample8(t *testing.T) {
	runSample(t, []byte("AAA"), 2)
}

func TestSample9(t *testing.T) {
	runSample(t, []byte("AABB"), 3)
}

func TestSample10(t *testing.T) {
	runSample(t, []byte("AABBAB"), 2)
}
