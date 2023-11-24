package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	var cnt int
	ask := func(arr []int) int {
		var res int
		for i := 0; i < len(arr); i++ {
			j := arr[i]
			res += a[j-1]
			if j == expect {
				res++
			}
			cnt++
			if cnt > 30 {
				t.Fatalf("Sample ask too much times %d", cnt)
			}
		}
		return res
	}

	res := solve(a, ask)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 5, 3, 4, 2}
	expect := 7
	runSample(t, a, expect)
}
