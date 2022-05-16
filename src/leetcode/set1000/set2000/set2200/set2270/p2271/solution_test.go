package p2271

import "testing"

func runSample(t *testing.T, tiles [][]int, carpetLen int, expect int) {
	res := maximumWhiteTiles(tiles, carpetLen)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tiles := [][]int{
		{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32},
	}
	carpetLen := 10
	expect := 9
	runSample(t, tiles, carpetLen, expect)
}

func TestSample2(t *testing.T) {
	tiles := [][]int{
		{10, 11}, {1, 1},
	}
	carpetLen := 2
	expect := 2
	runSample(t, tiles, carpetLen, expect)
}
