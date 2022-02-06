package main

import "testing"

func runSample(t *testing.T, n, k int, expect bool) {
	ok, res := solve(n, k)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if ok {
		buf := make([]byte, k+1)
		for i := 1; i <= k; i++ {
			buf[i] = '1'
		}
		for i, p := range res {
			l := 1 << uint(i)
			for j := p; j < p+l; j++ {
				if buf[j] == '1' {
					buf[j] = '0'
				} else {
					buf[j] = '1'
				}
			}
		}
		for i := 1; i <= k; i++ {
			if buf[i] == '1' {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 100, 89, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000, 898833, true)
}