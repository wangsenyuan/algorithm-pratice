package p3376

import "testing"

func runSample(t *testing.T, strength []int, k int, expect int) {
	res := findMinimumTime(strength, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []int{3, 4, 1}
	k := 1
	expect := 4
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := []int{2, 5, 4}
	k := 2
	expect := 5
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := []int{3}
	k := 8
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := []int{41, 42, 18, 46}
	k := 7
	expect := 29
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := []int{7, 3, 6, 18, 22, 50}
	k := 4
	expect := 12
	runSample(t, s, k, expect)
}
