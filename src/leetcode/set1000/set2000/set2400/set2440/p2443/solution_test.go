package p2443

import "testing"

func runSample(t *testing.T, arr []int, x int, y int, expect int64) {
	res := countSubarrays(arr, x, y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 3, 5, 2, 7, 5}
	x := 1
	y := 5
	var expect int64 = 2
	runSample(t, nums, x, y, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	x := 1
	y := 1
	var expect int64 = 10
	runSample(t, nums, x, y, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{928799, 888361, 928799, 928799, 928799, 928799, 124173, 93094, 399240, 946505, 93094, 93094, 585816}
	x := 93094
	y := 928799
	var expect int64 = 12
	runSample(t, nums, x, y, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{89992, 89992, 89992, 900911, 142432, 900911, 900911, 900911, 823426, 900911, 900911, 308091, 312853, 900911, 900911, 764677, 756995, 89992, 89992, 188452, 541874, 598970, 900911, 89992, 152245, 193942, 900911, 89992, 900911, 486074, 508973, 900911, 235617, 44768, 640310, 926517, 900911, 900911, 489462, 241420, 89992, 339246, 89992, 7549, 292723, 330338, 986407, 900911, 89992, 900911, 924927, 89992, 354279}
	x := 89992
	y := 900911
	var expect int64 = 405
	runSample(t, nums, x, y, expect)
}
