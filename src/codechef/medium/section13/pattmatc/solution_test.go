package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, S, T string, expect []int) {
	res := solve(S, T)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "*a*"
	T := "abacaba"
	expect := []int{1, 3, 3, 5, 5, 7, 7}
	runSample(t, S, T, expect)
}

func TestSample2(t *testing.T) {
	S := "*a*b*"
	T := "abacaba"
	expect := []int{2, 6, 6, 6, 6, -1, -1}
	runSample(t, S, T, expect)
}
