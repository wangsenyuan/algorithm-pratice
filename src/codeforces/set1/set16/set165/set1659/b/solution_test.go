package main

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res, f := solve(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	if check(s, f) != expect || sum(f) != k {
		t.Fatalf("Sample result not correct")
	}
}

func sum(arr []int) int {
	var res int
	for _, x := range arr {
		res += x
	}
	return res
}

func check(s string, f []int) string {
	n := len(s)
	flips := make([]int, n+1)

	for i, x := range f {
		if i > 0 {
			flips[0] += x
			flips[i] -= x
		}
		if i+1 < n {
			flips[i+1] += x
		}
	}

	buf := []byte(s)

	for i := 0; i < n; i++ {
		if i > 0 {
			flips[i] += flips[i-1]
		}
		if flips[i]&1 == 1 {
			if buf[i] == '0' {
				buf[i] = '1'
			} else {
				buf[i] = '0'
			}
		}
	}

	return string(buf)
}

func TestSample1(t *testing.T) {
	k := 3
	s := "100001"
	expect := "111110"
	runSample(t, s, k, expect)
}
