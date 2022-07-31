package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, k int, expect []int) {
	res := solve(A, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func check(arr []int, k int) bool {
	for i := 0; i < len(arr); i += k {
		cnt := make(map[int]int)
		for j := i; j < i+k; j++ {
			cnt[arr[j]]++
		}
		if len(cnt) != k {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	A := []int{4, 7, 7, 7}
	k := 2
	runSample(t, A, k, nil)
}

func TestSample2(t *testing.T) {
	A := []int{2, 10, 2, 8, 3, 6}
	k := 2
	expect := []int{2, 3, 2, 6, 8, 10}
	runSample(t, A, k, expect)
}
