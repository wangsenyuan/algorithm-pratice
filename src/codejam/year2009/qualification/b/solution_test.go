package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, H, W int, M [][]int, expect [][]byte) {
	res := solve(H, W, M)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", H, W, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	H, W := 3, 3
	M := [][]int{
		{9, 6, 3},
		{5, 9, 6},
		{3, 5, 9},
	}
	expect := [][]byte{
		[]byte("abb"),
		[]byte("aab"),
		[]byte("aaa"),
	}
	runSample(t, H, W, M, expect)
}

func TestSample2(t *testing.T) {
	H, W := 1, 10
	M := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 7},
	}
	expect := [][]byte{
		[]byte("aaaaaaaaab"),
	}
	runSample(t, H, W, M, expect)
}

func TestSample3(t *testing.T) {
	H, W := 2, 3
	M := [][]int{
		{7, 6, 7},
		{7, 6, 7},
	}
	expect := [][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
	}
	runSample(t, H, W, M, expect)
}

func TestSample4(t *testing.T) {
	H, W := 5, 5
	M := [][]int{
		{1, 2, 3, 4, 5},
		{2, 9, 3, 9, 6},
		{3, 3, 0, 8, 7},
		{4, 9, 8, 9, 8},
		{5, 6, 7, 8, 9},
	}
	expect := [][]byte{
		[]byte("aaaaa"),
		[]byte("aabba"),
		[]byte("abbba"),
		[]byte("abbba"),
		[]byte("aaaaa"),
	}
	runSample(t, H, W, M, expect)
}

func TestSample5(t *testing.T) {
	H, W := 2, 13
	M := [][]int{
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
	}
	expect := [][]byte{
		[]byte("abcdefghijklm"),
		[]byte("nopqrstuvwxyz"),
	}
	runSample(t, H, W, M, expect)
}
