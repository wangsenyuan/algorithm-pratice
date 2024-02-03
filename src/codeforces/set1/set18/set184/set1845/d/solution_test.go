package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	x := getMaxResult(a, res)
	y := getMaxResult(a, expect)

	if x != y {
		t.Fatalf("Sample expect %d, but got %d", y, x)
	}
}

func getMaxResult(a []int, k int) int {
	var sum int

	for _, num := range a {
		if sum < k {
			sum += num
		} else {
			// sum >= k
			if sum+num >= k {
				sum += num
			} else {
				sum = k
			}
		}
	}

	return sum
}

func TestSample1(t *testing.T) {
	a := []int{3, -2, 1, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, -2, -1}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 2}
	expect := 6
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 1, -3, 2, -1, -2, 2}
	expect := 6
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{4, -2}
	expect := 4
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{-2, 3, -3}
	expect := 1
	runSample(t, a, expect)
}
