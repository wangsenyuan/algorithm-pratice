package main

import "testing"

func runSample(t *testing.T, nums []string, expect int) {
	res := solve(nums)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []string{
		"ab",
		"de",
		"aj",
	}
	expect := 47
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []string{
		"abcdef",
		"ghij",
		"bdef",
		"accbd",
		"g",
	}
	expect := 136542
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []string{
		"a",
		"b",
		"c",
		"d",
		"f",
		"g",
		"h",
		"i",
		"j",
	}
	// 8 * 9 / 2
	expect := 45
	runSample(t, nums, expect)
}
