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

	for i := 0; i+1 < len(a); i++ {
		if a[i] < a[i+1] {
			if res[i] >= res[i+1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		} else if a[i] > a[i+1] {
			if res[i] <= res[i+1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		} else {
			if res[i] == res[i+1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 4, 2, 2}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 5, 7, 8, 10, 3, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 3, 7, 9, 8, 8, 8, 8, 7, 7, 7, 7, 5, 3, 3, 3, 3, 8, 8}
	expect := true
	runSample(t, a, expect)
}
