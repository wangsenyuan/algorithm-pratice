package p2953

import "testing"

func runSample(t *testing.T, a []int, target int, expect int) {
	res := minimumAddedCoins(a, target)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 4, 10}
	target := 19
	expect := 2
	runSample(t, a, target, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 4, 10, 5, 7, 19}
	target := 19
	expect := 1
	runSample(t, a, target, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1}
	target := 20
	expect := 3
	runSample(t, a, target, expect)
}

func TestSample4(t *testing.T) {
	a := []int{15, 1, 12}
	target := 43
	expect := 4
	runSample(t, a, target, expect)
}
