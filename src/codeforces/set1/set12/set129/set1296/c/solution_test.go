package main

import "testing"

func runSample(t *testing.T, s string, expect []int) {
	res := solve(s)

	if len(expect) == 0 {
		if len(res) > 0 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
		return
	}

	if len(res) == 0 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if expect[1]-expect[0]+1 != res[1]-res[0]+1 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "LRUD"
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "LURD"
	expect := []int{1, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "RRUDU"
	expect := []int{3, 4}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "LLDDR"
	runSample(t, s, nil)
}
