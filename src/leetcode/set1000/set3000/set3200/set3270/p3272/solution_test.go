package p3272

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := countGoodIntegers(n, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 27)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 4, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 6, 2468)
}

func TestSample4(t *testing.T) {
	runSample(t, 4, 6, 58)
}

func TestSample5(t *testing.T) {
	runSample(t, 10, 6, 13249798)
}
