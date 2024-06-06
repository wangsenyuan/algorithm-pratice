package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, p []int, expect []int) {
	res := solve(p)

	expect_score := getScore(p, expect)
	res_score := getScore(p, res)

	if expect_score != res_score {
		t.Fatalf("Sample expect %d, but got %d", expect_score, res_score)
	}

	sort.Ints(res)

	for i := 0; i < len(res); i++ {
		if res[i] != i+1 {
			t.Fatalf("Sample result %v, not a permutatiion", res)
		}
	}
}



func TestSample1(t *testing.T) {
	p := []int{1, 2, 3, 4}
	expect := []int{2, 4, 1, 3}
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{4, 3, 1, 2}
	expect := []int{3, 1, 4, 2}
	runSample(t, p, expect)
}

func TestSample3(t *testing.T) {
	p := []int{6, 5, 1, 4, 2, 3}
	expect := []int{2, 5, 1, 4, 3, 6}
	runSample(t, p, expect)
}

func TestSample4(t *testing.T) {
	p := []int{1, 2, 4, 5, 7, 6, 8, 3}
	expect := []int{5, 4, 8, 2, 7, 1, 6, 3}
	runSample(t, p, expect)
}

func TestSample5(t *testing.T) {
	p := []int{2, 1, 3, 4}
	expect := []int{2, 4, 1, 3}
	runSample(t, p, expect)
}

func TestSample6(t *testing.T) {
	p := []int{2, 1, 4, 3}
	expect := []int{3, 4, 2, 1}
	runSample(t, p, expect)
}
