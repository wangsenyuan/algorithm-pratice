package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d expect %v, but got %v", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []int{2, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, 3, []int{2, 3, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, 5, []int{2, 1, 4, 5, 3})
}

func TestSample4(t *testing.T) {
	runSample(t, 6, []int{2, 1, 4, 3, 6, 5})
}
