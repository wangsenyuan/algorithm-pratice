package main

import "testing"

func runSample(t *testing.T, n, m int, B []string, expect string) {
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		b[i] = []byte(B[i])
	}
	res := solve(n, m, b)

	if string(res) != expect {
		t.Errorf("Sample expect %s, but got %s", expect, string(res))
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	B := []string{
		"xab",
		"a#z",
		"caa",
	}
	expect := "xabza"
	runSample(t, n, m, B, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 4
	B := []string{
		"pyqs",
		"vcot",
		"qbiu",
		"lihj",
		"uvmz",
	}
	expect := "pvcbihjz"
	runSample(t, n, m, B, expect)
}
