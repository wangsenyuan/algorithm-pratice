package main

import "testing"

func runSample(t *testing.T, N int, digits string, M int, X int, Y int, expect int64) {
	res := solve(N, []byte(digits), M, X, Y)
	if res != expect {
		t.Errorf("sample %d %s %d %d %d, expect %d, but got %d", N, digits, M, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 3
	digits := "474"
	M, X, Y := 2, 1, 1
	runSample(t, N, digits, M, X, Y, 2)
}

//
//func TestSample2(t *testing.T) {
//	N := 34
//	digits := "6311861109697810998905373107116111"
//	M, X, Y := 10, 4, 25
//	runSample(t, N, digits, M, X, Y, 1)
//}
