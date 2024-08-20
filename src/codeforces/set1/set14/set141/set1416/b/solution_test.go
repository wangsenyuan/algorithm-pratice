package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	y := make([]int, len(a))
	copy(y, a)
	ok, res := solve(a)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}

	for _, op := range res {
		i, j, x := op[0], op[1], op[2]
		y[i-1] -= i * x
		y[j-1] += i * x
		if y[i-1] < 0 || y[j-1] < 0 {
			t.Fatalf("Sample result %v got a negative value", res)
		}
	}

	for i := 1; i < len(y); i++ {
		if y[i] != y[0] {
			t.Fatalf("Sample result %v got non-equal result %v", res, y)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 16, 4, 18}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	expect := false
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{11, 19, 1, 1, 3}
	expect := true
	runSample(t, a, expect)
}
func TestSample4(t *testing.T) {
	a := []int{6, 12, 3, 14, 5, 4, 3, 15, 15, 8, 11, 6, 4, 18, 2, 6, 8, 24, 14, 2, 8, 32}
	expect := true
	runSample(t, a, expect)
}
