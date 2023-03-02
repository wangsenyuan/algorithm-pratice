package main

import "testing"

func runSample(t *testing.T, n int, m int, S []int64) {
	res := solve(n, m, S)

	for i := 0; i < len(S); i++ {
		tmp := getValue(m, res[i])
		if tmp != S[i] {
			t.Fatalf("Sample result %v ( => %d), not correct, expect %d", res[i], tmp, S[i])
		}
	}
}

func getValue(m int, arr [][]int) int64 {
	p := arr[0][0]
	k := arr[0][1]
	var sum int64

	for i := 0; i < k; i++ {
		l := arr[1][i]
		if l <= 0 || l > m {
			return -1
		}
		sum += int64(p+i) * int64(l)
	}
	return sum
}

func TestSample1(t *testing.T) {
	n, m := 5, 4
	S := []int64{20, 36, 1, 60}
	runSample(t, n, m, S)
}
