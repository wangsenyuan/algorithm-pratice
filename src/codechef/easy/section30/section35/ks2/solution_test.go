package main

import "testing"

func runSample(t *testing.T, k int64, expect uint64) {
	res := solve(k)

	if res != expect {
		t.Errorf("Sample expect %d th %d, but got %d", k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 28)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 37)
}

func TestSample3(t *testing.T) {
	id := 1
	for i := 19; i < 10000; i++ {
		j := i
		var sum int
		for j > 0 {
			sum += (j % 10)
			j /= 10
		}
		sum = sum % 10

		if sum == 0 {
			runSample(t, int64(id), uint64(i))
			id++
		}
	}
}
