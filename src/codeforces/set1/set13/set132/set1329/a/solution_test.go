package main

import "testing"

func runSample(t *testing.T, n int, l []int, expect bool) {
	res := solve(n, l)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	arr := make([]int, n)
	for i, p := range res {
		p--
		for j := p; j < p+l[i]; j++ {
			if j == n {
				t.Fatalf("Sample result %v, starting position %d, not having enought distance from end, expecting %d", res, p, l[i])
			}
			arr[j] = i + 1
		}
	}
	m := len(l)

	cnt := make([]int, m+1)

	for i, x := range arr {
		if x == 0 {
			t.Fatalf("Sample result %v, creating arr %v gap at %d", res, arr, i)
		}
		cnt[x]++
	}

	for i := 1; i <= m; i++ {
		if cnt[i] == 0 {
			t.Fatalf("Sample result %v, no color found for %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	n := 5
	l := []int{3, 2, 2}
	expect := true
	runSample(t, n, l, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	l := []int{1}
	expect := false
	runSample(t, n, l, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	l := []int{1, 2}
	expect := false
	runSample(t, n, l, expect)
}

func TestSample4(t *testing.T) {
	n := 200
	l := []int{49, 35, 42, 47, 134, 118, 14, 148, 58, 159, 33, 33, 8, 123, 99, 126, 75, 94, 1, 141, 61, 79, 122, 31, 48, 7, 66, 97, 141, 43, 25, 141, 7, 56, 120, 55, 49, 37, 154, 56, 13, 59, 153, 133, 18, 1, 141, 24, 151, 125}
	expect := true
	runSample(t, n, l, expect)
}
