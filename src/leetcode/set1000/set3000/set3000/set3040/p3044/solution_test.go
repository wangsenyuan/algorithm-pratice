package p3044

import "testing"

func runSample(t *testing.T, mat [][]int, expect int) {
	res := mostFrequentPrime(mat)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 1},
		{9, 9},
		{1, 1},
	}
	expect := 19
	runSample(t, mat, expect)
}
