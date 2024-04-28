package main

import "testing"

func runSample(t *testing.T, x1, x2 int, c []int, expect bool) {
	res := solve(x1, x2, c)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	check := func(x int, arr []int) bool {
		var sum int
		n := len(arr)
		if n == 0 {
			return false
		}

		for _, i := range arr {
			if c[i-1]*n < x {
				return false
			}
			sum += c[i-1]
		}
		return sum >= x
	}

	if !check(x1, res[0]) || !check(x2, res[1]) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	x1, x2 := 8, 16
	c := []int{3, 5, 2, 9, 8, 7}
	expect := true
	runSample(t, x1, x2, c, expect)
}

func TestSample2(t *testing.T) {
	x1, x2 := 20, 32
	c := []int{21, 11, 11, 12}
	expect := true
	runSample(t, x1, x2, c, expect)
}

func TestSample3(t *testing.T) {
	x1, x2 := 11, 32
	c := []int{5, 5, 16, 16}
	expect := false
	runSample(t, x1, x2, c, expect)
}

func TestSample4(t *testing.T) {
	x1, x2 := 12, 20
	c := []int{7, 8, 4, 11, 9}
	expect := false
	runSample(t, x1, x2, c, expect)
}
