package main

import (
	"testing"
)

func runSample(t *testing.T, n int, vb int, vs int, X []int, exam []int, expect int) {
	res := solve(n, vb, vs, X, exam)

	if expect != res {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, vb, vs := 4, 5, 2
	X := []int{0, 2, 4, 6}
	room := []int{4, 1}
	expect := 3
	runSample(t, n, vb, vs, X, room, expect)
}

func TestSample2(t *testing.T) {
	n, vb, vs := 2, 1, 1
	X := []int{0, 100000}
	room := []int{100000, 100000}
	expect := 2
	runSample(t, n, vb, vs, X, room, expect)
}

func TestSample3(t *testing.T) {
	n, vb, vs := 35, 35, 12
	X := []int{0, 90486, 90543, 90763, 91127, 91310, 92047, 92405, 93654, 93814, 94633, 94752, 94969, 94994, 95287, 96349, 96362, 96723, 96855, 96883, 97470, 97482, 97683, 97844, 97926, 98437, 98724, 98899, 98921, 99230, 99253, 99328, 99444, 99691, 99947}
	room := []int{96233, -7777}
	expect := 9
	runSample(t, n, vb, vs, X, room, expect)
}
