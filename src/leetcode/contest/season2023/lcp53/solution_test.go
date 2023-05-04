package lcp53

import "testing"

func runSample(t *testing.T, time []int, position []int, expect int) {
	res := defendSpaceCity(time, position)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	time := []int{1, 2, 1}
	pos := []int{6, 3, 3}
	expect := 5
	runSample(t, time, pos, expect)
}

func TestSample2(t *testing.T) {
	time := []int{1, 1, 1, 2, 2, 3, 5}
	pos := []int{1, 2, 3, 1, 2, 1, 3}
	expect := 9
	runSample(t, time, pos, expect)
}

func TestSample3(t *testing.T) {
	time := []int{1, 4, 3, 2, 5, 3, 2, 1, 1, 2, 5, 5, 3, 5, 4, 1, 1, 2, 4, 5, 5, 5, 1, 3, 4, 4, 4, 4, 2, 3, 2, 1, 5, 5, 2, 1, 4, 1, 4, 5, 4, 4, 2, 3, 4, 5, 5, 3, 3, 2, 1, 5, 5}
	pos := []int{4, 29, 20, 12, 11, 7, 15, 17, 5, 25, 21, 15, 5, 3, 10, 10, 29, 11, 3, 25, 28, 14, 26, 19, 9, 25, 28, 12, 4, 22, 19, 23, 23, 13, 6, 3, 23, 11, 1, 26, 15, 27, 5, 16, 7, 4, 24, 23, 2, 20, 8, 20, 27}
	expect := 75
	runSample(t, time, pos, expect)
}
