package main

import "testing"

func runSample(t *testing.T, A string, expect int) {
	B, C := solve(A)
	res := count(B) + count(C)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func count(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res += int(s[i] - '0')
	}
	return res
}

func TestSample1(t *testing.T) {
	A := "00111"
	expect := 2
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := "10"
	expect := 1
	runSample(t, A, expect)
}
