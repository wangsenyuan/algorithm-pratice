package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	if !expect {
		return
	}

	for i := 0; i+1 < len(res); i++ {
		if a[i] > 0 && res[i] != a[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		if res[i]/2 != res[i+1] && res[i+1]/2 != res[i] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{-1, -1, -1, 2, -1, -1, 1, -1}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, -1, -1, -1}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, -1, -1, -1, 9, -1}
	expect := true
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-1, 5, -1, 6}
	expect := false
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, -1, -1, 3}
	expect := false
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{1, 2, 3, 4}
	expect := false
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{4, 2}
	expect := true
	runSample(t, a, expect)
}

func TestSample8(t *testing.T) {
	a := []int{-1, 3, -1, 3, 6}
	expect := true
	runSample(t, a, expect)
}

func TestSample9(t *testing.T) {
	a := []int{-1, -1, 3, -1, -1, -1, -1, 7, -1, -1, 3, -1, -1}
	expect := true
	runSample(t, a, expect)
}
