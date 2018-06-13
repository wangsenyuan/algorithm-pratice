package main

import (
	"reflect"
	"testing"
)

func TestSolve1(t *testing.T) {
	n := 5
	chefs := []Chef{Chef{2, 3}, Chef{1, 7}, Chef{5, 5}, Chef{3, 1}, Chef{8, 15}}
	res := solve(n, chefs)
	expect := []int64{3, 4, 5, 4, 9}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("expect %v, but got %v", expect, res)
	}
}
