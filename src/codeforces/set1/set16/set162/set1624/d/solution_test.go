package main

import "testing"

func runSample(t *testing.T, n int, k int, s string, expect int) {
	res := solve(n, k, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 8, 2
	s := "bxyaxzay"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample2(t *testing.T) {
	n, k := 6, 3
	s := "aaaaaa"
	expect := 2
	runSample(t, n, k, s, expect)
}

func TestSample3(t *testing.T) {
	n, k := 6, 1
	s := "abcdef"
	expect := 1
	runSample(t, n, k, s, expect)
}

func TestSample4(t *testing.T) {
	n, k := 6, 6
	s := "abcdef"
	expect := 1
	runSample(t, n, k, s, expect)
}
func TestSample5(t *testing.T) {
	n, k := 3, 2
	s := "dxd"
	expect := 1
	runSample(t, n, k, s, expect)
}


func TestSample6(t *testing.T) {
	n, k := 11, 2
	s := "abcabcabcac"
	expect := 5
	runSample(t, n, k, s, expect)
}

func TestSample7(t *testing.T) {
	n, k := 6, 6
	s := "sipkic"
	expect := 1
	runSample(t, n, k, s, expect)
}

func TestSample8(t *testing.T) {
	n, k := 7, 2
	s := "eatoohd"
	expect := 1
	runSample(t, n, k, s, expect)
}


func TestSample9(t *testing.T) {
	n, k := 3, 1
	s := "llw"
	expect := 3
	runSample(t, n, k, s, expect)
}

func TestSample10(t *testing.T) {
	n, k := 6, 2
	s := "bfvfbv"
	expect := 3
	runSample(t, n, k, s, expect)
}