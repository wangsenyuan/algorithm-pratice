package main

import "testing"

func runSample(t *testing.T, n int, A, B string, P int, S int, cnt int) {
	sz, tcnt := solve(n, A, B, P)
	if sz != S || tcnt != cnt {
		t.Errorf("this sample %d, %s, %s, %d should give answer %d %d, but got %d %d", n, A, B, P, S, cnt, sz, tcnt)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, "abcdefgh", "abcdefgh", 40, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, "abcdefgh", "ccccccch", 100, 1, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, "a", "b", 100, 0, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, "a", "b", 0, 1, 1)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, "aaaaaaaa", "bbbbbbbb", 0, 2, 1)
}
