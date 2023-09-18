package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, b []int) {
	k, res := solve(b)
	// res is a
	n := len(b)
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		x := res[i-1]
		if x <= k {
			c[x] = n + 1
			for j := i - 1; j > 0; j-- {
				if res[j-1] > k {
					c[x] = res[j-1]
					break
				}
			}
		} else {
			// x < k
			c[x] = 0
			for j := i - 1; j > 0; j-- {
				if res[j-1] <= k {
					c[x] = res[j-1]
					break
				}
			}
		}
	}

	if !reflect.DeepEqual(c[1:], b) {
		t.Fatalf("Sample result %d %v, not correct, it gots %v", k, res, c[1:])
	}
}

func TestSample1(t *testing.T) {
	b := []int{5, 3, 1, 2}
	runSample(t, b)
}

func TestSample2(t *testing.T) {
	b := []int{7, 7, 7, 3, 3, 3}
	runSample(t, b)
}

func TestSample3(t *testing.T) {
	b := []int{4, 4, 4, 0, 0, 0}
	runSample(t, b)
}
