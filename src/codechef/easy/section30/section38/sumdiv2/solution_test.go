package main

import "testing"

func runSample(t *testing.T, x, y int) {
	N := solve(int64(x), int64(y))

	if N < 1 {
		t.Fatalf("Sample result %d, not correct", N)
	}

	n := int(N)

	if (n+y)%x != 0 {
		t.Fatalf("Sample result (%d + %d) div %d != 0", n, y, x)
	}

	if (n+x)%y != 0 {
		t.Fatalf("Sample result (%d + %d) div %d != 0", n, x, y)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 18, 42)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1)
}
func TestSample3(t *testing.T) {
	runSample(t, 100, 200)
}

func TestSample4(t *testing.T) {
	runSample(t, 150, 1500)
}
