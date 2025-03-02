package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect string) {
	res := solve(n, m, k)
	if res == expect {
		return
	}
	if res == "-1" || expect == "-1" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	var best int

	for i := 0; i < n+m; i++ {
		cnt := []int{0, 0}
		for j := i; j < n+m; j++ {
			cnt[int(res[j]-'0')]++
			tmp := abs(cnt[0] - cnt[1])
			best = max(best, tmp)
		}
	}

	if best != k {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 1, 2, 1
	expect := "101"
	runSample(t, n, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 4, 3, 2
	expect := "0100101"
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 2, 4, 3
	expect := "011011"
	runSample(t, n, m, k, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 8, 3, 2
	expect := "-1"
	runSample(t, n, m, k, expect)
}

func TestSample5(t *testing.T) {
	n, m, k := 5, 0, 4
	expect := "-1"
	runSample(t, n, m, k, expect)
}

func TestSample6(t *testing.T) {
	n, m, k := 5, 0, 5
	expect := "00000"
	runSample(t, n, m, k, expect)
}

func TestSample7(t *testing.T) {
	n, m, k := 26, 25, 11
	expect := "000000000001010101010101010101010101010101111111111"
	runSample(t, n, m, k, expect)
}
