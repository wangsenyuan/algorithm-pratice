package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, L []uint64, A, B, C uint64, S string, expect []uint64) {
	res := solve(n, L, A, B, C, []byte(S))
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v %d %d %d %s, expect %v, but got %v", n, L, A, B, C, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var A, B, C uint64 = 2, 3, 1000
	L := []uint64{1, 1, 1}
	S := "ARM"
	expect := []uint64{3, 3, 9}
	runSample(t, n, L, A, B, C, S, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	var A, B, C uint64 = 0, 1, 1000
	L := []uint64{1, 2, 3, 4}
	S := "AMAM"
	expect := []uint64{1, 2, 3, 4}
	runSample(t, n, L, A, B, C, S, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	var A, B, C uint64 = 0, 2, 1000
	L := []uint64{1, 2, 3, 4}
	S := "AMAM"
	expect := []uint64{1, 4, 6, 16}
	runSample(t, n, L, A, B, C, S, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	var A, B, C uint64 = 1, 2, 1000
	L := []uint64{1, 2, 3, 4}
	S := "AMAM"
	expect := []uint64{2, 6, 9, 22}
	runSample(t, n, L, A, B, C, S, expect)
}

func TestSample6(t *testing.T) {
	n := 6
	var A, B, C uint64 = 1, 2, 1000
	L := []uint64{1, 2, 3, 4, 5, 6}
	S := "AMAMRA"
	expect := []uint64{2, 6, 9, 22, 30, 27}
	runSample(t, n, L, A, B, C, S, expect)
}
