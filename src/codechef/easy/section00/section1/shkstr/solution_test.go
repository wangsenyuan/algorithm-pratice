package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, S []string, Q int, R []int, P []string, expect []string) {
	SS := make([][]byte, n)
	for i := 0; i < n; i++ {
		SS[i] = []byte(S[i])
	}

	PP := make([][]byte, Q)
	for i := 0; i < Q; i++ {
		PP[i] = []byte(P[i])
	}

	res := solve(n, SS, Q, R, PP)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %d %v %v, expect %v, but got %v", n, S, Q, R, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	S := []string{
		"abcd",
		"abce",
		"abcdex",
		"abcde",
	}
	Q := 3
	R := []int{3, 3, 4}
	P := []string{"abcy", "abcde", "abcde"}
	expect := []string{"abcd", "abcdex", "abcde"}
	runSample(t, n, S, Q, R, P, expect)
}
