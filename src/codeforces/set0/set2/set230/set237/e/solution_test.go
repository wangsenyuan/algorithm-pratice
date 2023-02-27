package main

import "testing"

func runSample(t *testing.T, T string, S []string, A []int, expect int64) {
	res := solve(T, S, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	T := "bbaze"
	S := []string{
		"bzb",
		"aeb",
		"ba",
	}
	A := []int{2, 3, 10}
	var expect int64 = 8
	runSample(t, T, S, A, expect)
}

func TestSample2(t *testing.T) {
	T := "abacaba"
	S := []string{
		"aba",
		"bcc",
		"caa",
		"bbb",
	}
	A := []int{2, 1, 2, 5}
	var expect int64 = 18
	runSample(t, T, S, A, expect)
}

func TestSample3(t *testing.T) {
	T := "xyz"
	S := []string{
		"axx",
		"za",
		"efg",
		"t",
	}
	A := []int{8, 1, 4, 1}
	var expect int64 = -1
	runSample(t, T, S, A, expect)
}
