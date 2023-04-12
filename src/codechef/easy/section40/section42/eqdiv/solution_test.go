package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, nums []int, expect []int) {
	res, ass := solve(k, nums)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	for i := 0; i < len(nums); i++ {
		check(t, expect[i], k, nums[i], ass[i])
	}
}

func check(t *testing.T, expect int, k int, n int, res string) {
	var a int64
	var b int64

	for i := 0; i < n; i++ {
		if res[i] == '1' {
			a += int64(pow(i+1, k))
		} else {
			b += int64(pow(i+1, k))
		}
	}
	a -= b
	if a < 0 {
		a *= -1
	}
	if a != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, a)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	nums := []int{4, 5}
	expect := []int{2, 3}
	runSample(t, k, nums, expect)
}

func TestSample2(t *testing.T) {
	k := 4
	nums := []int{5, 6, 9}
	expect := []int{271, 317, 253}
	runSample(t, k, nums, expect)
}
