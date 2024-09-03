package main

import "testing"

func runSample(t *testing.T, nums []string, k int, expect string) {
	res := solve(nums, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []string{
		"0000000",
	}
	k := 7
	expect := "8"
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []string{
		"0010010",
		"0010010",
	}
	k := 5
	expect := "97"
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []string{
		"0100001",
		"1001001",
		"1010011",
	}
	k := 5
	expect := ""
	runSample(t, nums, k, expect)
}
