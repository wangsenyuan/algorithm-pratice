package p3020

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := flowerGame(n, m)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 0)
}
