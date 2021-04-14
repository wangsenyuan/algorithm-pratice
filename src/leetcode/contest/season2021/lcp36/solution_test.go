package lcp36

import "testing"

func runSample(t *testing.T, tiles []int, expect int) {
	res := maxGroupNumber(tiles)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tiles := []int{2, 2, 2, 3, 4}
	expect := 1
	runSample(t, tiles, expect)
}

func TestSample2(t *testing.T) {
	tiles := []int{2, 2, 2, 3, 4, 1, 3}
	expect := 2
	runSample(t, tiles, expect)
}

func TestSample3(t *testing.T) {
	tiles := []int{1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 10, 10, 10, 10, 11, 11, 11, 12}
	expect := 13
	runSample(t, tiles, expect)
}
