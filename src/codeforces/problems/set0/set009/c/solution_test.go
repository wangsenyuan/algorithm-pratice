package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 2)
}

func TestSample2(t *testing.T) {
	var cnt int

	for i := 1; i <= 10000; i++ {
		cnt += check(i)
		runSample(t, i, cnt)
	}
}

func check(num int) int {
	for num > 0 {
		x := num % 10
		if x > 1 {
			return 0
		}
		num /= 10
	}
	return 1
}
