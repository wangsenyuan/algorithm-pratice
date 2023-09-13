package main

import "testing"

func runSample(t *testing.T, L []int, R []int, expect int) {
	res := solve(L, R)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := []int{1, 3, 3, 3, 3}
	R := []int{3, 3, 2, 2, 2}
	expect := 4
	runSample(t, L, R, expect)
}

func TestSample2(t *testing.T) {
	L := []int{1, 3776, 3776, 8848, 8848, 8848, 8848, 8848, 8848, 8848}
	R := []int{8848, 8848, 8848, 8848, 8848, 8848, 8848, 8848, 3776, 5}
	expect := 884111967
	runSample(t, L, R, expect)
}
