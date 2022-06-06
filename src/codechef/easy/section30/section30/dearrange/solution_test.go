package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if len(res) != expect {
		t.Fatalf("Sample expect %d steps, but got %v", expect, res)
	}

	if expect > 0 {
		var B []int
		for _, cur := range res {
			l, r := cur.l, cur.r
			B = cur.P
			cnt := make(map[int]int)
			for i := l - 1; i <= r-1; i++ {
				cnt[A[i]]++
				cnt[B[i]]--
				if B[i] == A[i] {
					t.Fatalf("Sample result not correct")
				}
			}
			for _, v := range cnt {
				if v != 0 {
					t.Fatalf("Sample result %v, not correct", res)
				}
			}
			A = B
		}

		if !sort.IntsAreSorted(B) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 1}
	runSample(t, A, 2)
}

func TestSample2(t *testing.T) {
	A := []int{2, 1, 3, 5, 4}
	runSample(t, A, 2)
}
