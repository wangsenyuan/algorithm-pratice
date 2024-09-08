package p3279

import "testing"

func runSample(t *testing.T, start []int, d int, expect int) {
	res := maxPossibleScore(start, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	start := []int{6, 0, 3}
	d := 2
	expect := 4
	runSample(t, start, d, expect)
}

func TestSample2(t *testing.T) {
	start := []int{2, 6, 13, 13}
	d := 5
	expect := 5
	runSample(t, start, d, expect)
}

func TestSample3(t *testing.T) {
	start := []int{1000000000, 0}
	d := 1000000000
	expect := 2000000000
	runSample(t, start, d, expect)
}

func TestSample4(t *testing.T) {
	start := []int{0, 9, 2, 9}
	d := 2
	expect := 2
	runSample(t, start, d, expect)
}

