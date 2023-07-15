package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cnt []int, expect []string) {
	res := solve(cnt)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cnt := []int{1, 2, 0, 1}
	expect := []string{"Ya", "Ya", "Tidak", "Tidak"}
	runSample(t, cnt, expect)
}

func TestSample2(t *testing.T) {
	cnt := []int{0, 1, 0, 0}
	expect := []string{"Tidak", "Ya", "Tidak", "Tidak"}
	runSample(t, cnt, expect)
}
