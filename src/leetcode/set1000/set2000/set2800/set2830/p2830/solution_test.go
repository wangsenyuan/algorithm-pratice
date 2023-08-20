package p2830

import "testing"

func runSample(t *testing.T, n int, offers [][]int, expect int) {
	res := maximizeTheProfit(n, offers)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	offers := [][]int{
		{0, 0, 1}, {0, 2, 2}, {1, 3, 2},
	}
	expect := 3
	runSample(t, n, offers, expect)
}
