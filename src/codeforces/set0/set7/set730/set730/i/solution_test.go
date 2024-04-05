package main

import "testing"

func runSample(t *testing.T, a []int, b []int, p int, s int, expect int) {
	gain, first, second := solve(a, b, p, s)

	if len(first) != p || len(second) != s {
		t.Fatalf("Sample result %v %v not having correct length", first, second)
	}

	if gain != expect {
		t.Fatalf("Sample result not get the max answer, expect %d, but got %d", expect, gain)
	}

	var sum int
	for _, x := range first {
		sum += a[x-1]
	}

	for _, y := range second {
		sum += b[y-1]
	}

	if sum != expect {
		t.Fatalf("Sample result not giving the correct answer, expect %d, but got %d", expect, sum)
	}
}

func TestSample1(t *testing.T) {
	p, s := 2, 2
	a := []int{1, 3, 4, 5, 2}
	b := []int{5, 3, 2, 1, 4}
	expect := 18
	runSample(t, a, b, p, s, expect)
}

func TestSample2(t *testing.T) {
	p, s := 2, 2
	a := []int{10, 8, 8, 3}
	b := []int{10, 7, 9, 4}
	expect := 31
	runSample(t, a, b, p, s, expect)
}

func TestSample3(t *testing.T) {
	p, s := 3, 1
	a := []int{5, 2, 5, 1, 7}
	b := []int{6, 3, 1, 6, 3}
	expect := 23
	runSample(t, a, b, p, s, expect)
}

func TestSample4(t *testing.T) {
	p, s := 1, 2
	a := []int{4, 2, 4, 5}
	b := []int{3, 2, 5, 3}
	expect := 13
	runSample(t, a, b, p, s, expect)
}
