package main

import "testing"

func runSample(t *testing.T, n, d int64, expectAns int, expectCnt int) {
	resAns, resCnt := solve(n, d)

	if resAns != expectAns || resCnt != expectCnt {
		t.Errorf("Sample %d %d, expect %d %d, but got %d %d", n, d, expectAns, expectCnt, resAns, resCnt)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 1, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 3, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 11, 13, 1, 4)
}
