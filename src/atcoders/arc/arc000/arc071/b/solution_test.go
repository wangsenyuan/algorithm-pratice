package main

import "testing"

func runSample(t *testing.T, xs []int, ys []int, expect int) {
	res := solve(xs, ys)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	xs := []int{1, 3, 4}
	ys := []int{1, 3, 6}
	expect := 60
	runSample(t, xs, ys, expect)
}

func TestSample2(t *testing.T) {
	xs := []int{-790013317, -192321079, 95834122, 418379342, 586260100, 802780784}
	ys := []int{-253230108, 193944314, 363756450, 712662868, 735867677}
	expect := 835067060
	runSample(t, xs, ys, expect)
}
