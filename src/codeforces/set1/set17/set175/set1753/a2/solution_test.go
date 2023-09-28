package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	prev := 0
	var sum int
	for _, cur := range res {
		l, r := cur[0], cur[1]
		l--
		for i := l; i < r; i++ {
			if (i-l)&1 == 0 {
				sum += a[i]
			} else {
				sum -= a[i]
			}
		}
		if prev != l {
			t.Fatalf("Sample result %v, not cover whole array", res)
		}
		prev = r
	}
	if sum != 0 {
		t.Fatalf("Sample result %v, got wrong sum %d", res, sum)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, 1, 1, 1, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, -1, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{0, 0, 0, 0}
	expect := true
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{-1, 1, 0, 1, 0, 1, 0}
	expect := true
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{0, -1, 1, 0, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 0, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{0, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample9(t *testing.T) {
	a := []int{1}
	expect := false
	runSample(t, a, expect)
}
