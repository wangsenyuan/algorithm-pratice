package main

import "testing"

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	cnt1 := countCharacter(expect)
	cnt2 := countCharacter(res)

	if cnt1 != cnt2 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func countCharacter(arr []int) int {
	res := make(map[int]int)
	for i := 1; i < len(arr); i++ {
		res[arr[i]-arr[i-1]]++
	}

	return len(res)
}

func TestSample1(t *testing.T) {
	k, n := 5, 9
	expect := []int{1, 3, 4, 7, 8}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	k, n := 4, 12
	expect := []int{2, 4, 7, 12}
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	k, n := 3, 3
	expect := []int{1, 2, 3}
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	k, n := 3, 4
	expect := []int{1, 3, 4}
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	k, n := 4, 6
	expect := []int{2, 4, 5, 6}
	runSample(t, n, k, expect)
}

func TestSample6(t *testing.T) {
	k, n := 8, 11
	expect := []int{1, 2, 3, 5, 6, 7, 8, 11}
	runSample(t, n, k, expect)
}

func TestSample7(t *testing.T) {
	k, n := 6, 20
	expect := []int{1, 2, 4, 7, 11, 20}
	runSample(t, n, k, expect)
}
