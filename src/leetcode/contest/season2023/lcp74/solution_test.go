package lcp74

import "testing"

func runSample(t *testing.T, fields [][]int, expect int) {
	res := fieldOfGreatestBlessing(fields)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	fields := [][]int{
		{0, 0, 1}, {1, 0, 1},
	}
	expect := 2
	runSample(t, fields, expect)
}

func TestSample2(t *testing.T) {
	fields := [][]int{
		{4, 4, 6}, {7, 5, 3}, {1, 6, 2}, {5, 6, 3},
	}
	expect := 3
	runSample(t, fields, expect)
}
