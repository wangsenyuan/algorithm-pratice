package main

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	n := 5
	parent := []int{0, 1, 2, 3}
	club := []int{0, 0, 0, 0, 0}
	K := []int{2, 1, 1, 0, 0}
	expect := []int64{4, 2, 2, 1, 1}

	res := solve(n, parent, club, K)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("expect %v, but got %v", expect, res)
	}
}
