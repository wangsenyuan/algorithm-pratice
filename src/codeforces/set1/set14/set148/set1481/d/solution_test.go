package main

import "testing"

func runSample(t *testing.T, n, m int, G []string, expect bool) {
	g := make([][]byte, n)
	for i := 0; i < n; i++ {
		g[i] = []byte(G[i])
	}

	res := solve(n, m, g)

	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if expect {
		buf := make([]byte, m)
		for i := 1; i <= m; i++ {
			buf[i-1] = G[res[i-1]-1][res[i]-1]
		}

		for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
			if buf[i] != buf[j] {
				t.Fatalf("Sample result %v, got %s, not palindrome", res, string(buf))
			}
		}
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 1
	G := []string{
		"*ba",
		"b*b",
		"ab*",
	}
	expect := true

	runSample(t, n, m, G, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	G := []string{
		"*ba",
		"b*b",
		"ab*",
	}
	expect := true

	runSample(t, n, m, G, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 4
	G := []string{
		"*ba",
		"b*b",
		"ab*",
	}
	expect := true

	runSample(t, n, m, G, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 6
	G := []string{
		"*aaa",
		"b*ba",
		"ab*a",
		"bba*",
	}
	expect := true

	runSample(t, n, m, G, expect)
}

func TestSample5(t *testing.T) {
	n, m := 2, 6
	G := []string{
		"*a",
		"b*",
	}
	expect := false

	runSample(t, n, m, G, expect)
}
