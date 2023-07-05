package main

import "testing"

func runSample(t *testing.T, nums []int, expect string) {
	res := solve(nums)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	var arr []int
	n := len(nums)
	l, r := 0, n-1
	for i := 0; i < len(res); i++ {
		if res[i] == 'L' {
			arr = append(arr, nums[l])
			l++
		} else {
			arr = append(arr, nums[r])
			r--
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			t.Fatalf("Sample result %s, got wrong answer %v", res, arr)
		}
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 4, 3, 2}
	expect := "LRRR"
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 3, 5, 6, 5, 4, 2}
	expect := "LRLRRR"
	runSample(t, nums, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{2, 2, 2}
	expect := "R"
	runSample(t, nums, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{1, 2, 4, 3}
	expect := "LLRR"
	runSample(t, nums, expect)
}

func TestSample5(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	expect := "RRRRR"
	runSample(t, nums, expect)
}

func TestSample6(t *testing.T) {
	nums := []int{5, 4, 2, 2, 1}
	expect := "RRL"
	runSample(t, nums, expect)
}
