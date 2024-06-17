package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	n := len(a)

	b := make([]int, n)
	copy(b, a)

	res, steps := solve(b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}

	vis := make([]int, n+2)
	apply := func(l int, r int) {
		l--
		r--
		var mex int
		for i := l; i <= r; i++ {
			if a[i] <= n {
				vis[a[i]]++
			}
			for vis[mex] > 0 {
				mex++
			}
		}

		for i := l; i <= r; i++ {
			a[i] = mex
		}
		for i := 0; i <= n; i++ {
			vis[i] = 0
		}
	}

	for _, step := range steps {
		apply(step[0], step[1])
	}
	var sum int
	for i := 0; i < n; i++ {
		sum += a[i]
	}

	if sum != expect {
		t.Fatalf("Sample %v =>  got %d, not correct", res, sum)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 1}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 100, 2, 1}
	expect := 105
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0}
	expect := 1
	runSample(t, a, expect)
}
