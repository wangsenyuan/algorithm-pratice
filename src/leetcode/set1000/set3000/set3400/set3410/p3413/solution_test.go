package p3413

import "testing"

func runSample(t *testing.T, coins [][]int, k int, expect int) {
	res := maximumCoins(coins, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	coins := [][]int{
		{8, 10, 1}, {1, 3, 2}, {5, 6, 4},
	}
	k := 4
	expect := 10
	runSample(t, coins, k, expect)
}

func TestSample2(t *testing.T) {
	coins := [][]int{
		{1, 10, 3},
	}
	k := 2
	expect := 6
	runSample(t, coins, k, expect)
}

func TestSample3(t *testing.T) {
	coins := [][]int{
		{20, 27, 18}, {37, 40, 19}, {11, 14, 18}, {8, 10, 9}, {28, 32, 14}, {1, 7, 5},
	}
	k := 32
	expect := 380
	runSample(t, coins, k, expect)
}
