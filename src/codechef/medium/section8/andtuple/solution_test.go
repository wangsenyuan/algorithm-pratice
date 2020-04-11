package main

import "testing"

func runSample(t *testing.T, num uint64, k int, expect int) {
	res := solve(num, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", num, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 4, 2)
}

func TestSample3(t *testing.T) {
	for num := 2; num <= 10000; num++ {
		a := solve1(uint64(num), 3)
		b := solve(uint64(num), 3)
		if a != b {
			t.Fatalf("not match at %d, %d vs %d", num, a, b)
		}
	}
}

func TestSample4(t *testing.T) {
	for num := 2; num <= 10000; num++ {
		a := solve1(uint64(num), 4)
		b := solve(uint64(num), 4)
		if a != b {
			t.Fatalf("not match at %d, %d vs %d", num, a, b)
		}
	}
}
