package main

import "testing"

func runSample(t *testing.T, n int, m int, expect string) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, "Alice")
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 6, "Alice")
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 9, "Bob")
}

func TestSample4(t *testing.T) {
	runSample(t, 45, 58, "Bob")
}

func TestSample5(t *testing.T) {
	runSample(t, 39, 94, "Alice")
}

func TestSample6(t *testing.T) {
	runSample(t, 36, 54, "Bob")
}
