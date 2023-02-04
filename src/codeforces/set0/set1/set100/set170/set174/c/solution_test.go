package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	arr := make([]int, len(A))

	for _, op := range res {
		i, j := op[0], op[1]
		arr[i-1]++
		if j < len(A) {
			arr[j]--
		}
	}

	for i := 1; i < len(A); i++ {
		arr[i] += arr[i-1]
	}

	if !reflect.DeepEqual(arr, A) {
		t.Errorf("Sample result %v, got different array %v, expect %v", res, arr, A)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1, 1, 4, 1}
	expect := 5
	runSample(t, A, expect)
}
