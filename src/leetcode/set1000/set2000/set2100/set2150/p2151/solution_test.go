package p2151

import "testing"

func runSample(t *testing.T, statements [][]int, expect int) {
	res := maximumGood(statements)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	statements := [][]int{
		{2, 1, 2}, {1, 2, 2}, {2, 0, 2},
	}
	expect := 2
	runSample(t, statements, expect)
}

func TestSample2(t *testing.T) {
	statements := [][]int{
		{2, 0}, {0, 2},
	}
	expect := 1
	runSample(t, statements, expect)
}

