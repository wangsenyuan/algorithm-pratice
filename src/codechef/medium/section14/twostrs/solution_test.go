package main

import "testing"

func runSample(t *testing.T, A, B string, S []string, bb []int, expect int64) {
	res := solve(A, B, len(S), S, bb)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := "hitech"
	B := "codechef"
	S := []string{
		"chef",
		"code",
		"eche",
	}
	bb := []int{3, 1, 5}
	var expect int64 = 13

	runSample(t, A, B, S, bb, expect)
}

func TestSample2(t *testing.T) {
	A := "aaaaaaaaaaaaaaa"
	B := "bbbbaaaaaaaaaaa"
	S := []string{
		"aaaaaaaaaaaaaaaaaaaaaaaaaa",
	}
	bb := []int{200}
	var expect int64 = 200

	runSample(t, A, B, S, bb, expect)
}
