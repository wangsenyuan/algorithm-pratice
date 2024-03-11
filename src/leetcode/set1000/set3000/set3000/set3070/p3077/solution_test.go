package p3077

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := maximumStrength(a, k)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, -1, 2}
	k := 3
	expect := 22
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{12, -2, -2, -2, -2}
	k := 5
	expect := 64
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1, -2, -3}
	k := 1
	expect := -1
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{7, -70, 75}
	k := 1
	expect := 75
	runSample(t, a, k, expect)
}
