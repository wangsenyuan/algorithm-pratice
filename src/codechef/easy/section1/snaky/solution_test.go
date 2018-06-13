package main

import "testing"

func runSample(t *testing.T, n, m, x, y, l int, ds []byte, ew string, en int) {
	w, s := solve(n, m, x, y, l, ds)
	if w != ew {
		t.Errorf("the sample %v, snake should collides with %s, but got %s", ds, ew, w)
	}

	if s != en {
		t.Errorf("the sample %v, should collides after %d moves, but got %d", ds, en, s)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 10, 3, 2, 6, []byte("URDDL"), "WALL", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 6, 1, 6, 13, []byte("RRRRRDDDLLLU"), "BODY", 1)
}
