package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect bool) {
	res := solve(n, m, k)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	// 计算水平的
	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j+1 < m && res[i][j] == res[i][j+1] {
				cnt++
			}
		}
	}

	if cnt != k {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 2, 17, 16
	expect := true
	runSample(t, n, m, k, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 4, 4, 2
	expect := true
	runSample(t, n, m, k, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 10, 10, 24
	expect := true
	runSample(t, n, m, k, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 2, 3, 0
	expect := true
	runSample(t, n, m, k, expect)
}
func TestSample5(t *testing.T) {
	n, m, k := 3, 2, 3
	expect := true
	runSample(t, n, m, k, expect)
}

func TestSample6(t *testing.T) {
	n, m, k := 101, 34, 31
	expect := true
	runSample(t, n, m, k, expect)
}

func TestSample7(t *testing.T) {
	n, m, k := 34, 101, 32
	expect := true
	runSample(t, n, m, k, expect)
}