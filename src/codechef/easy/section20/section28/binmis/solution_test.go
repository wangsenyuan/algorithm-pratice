package main

import "testing"

func runSample(t *testing.T, n int, s string, expect bool) {
	res := solve(n, s)

	if len(res) > 0 != expect {
		t.Errorf("Sample expect %t, but got %v", expect, res)
	} else if expect {
		if !check(s, res) {
			t.Errorf("Sample result %v, not correct", res)
		}
	}
}

func check(s string, res []int) bool {
	l, r := res[0], res[1]
	l--
	r--
	cnt := make([]int, 2)
	n := len(s)
	for i := 0; i < n; i++ {
		if i >= l && i <= r {
			cnt[1-int(s[i]-'0')]++
		} else {
			cnt[int(s[i]-'0')]++
		}
	}
	return cnt[0] == cnt[1]
}

func TestSample1(t *testing.T) {
	n := 2
	s := "01"
	expect := true
	runSample(t, n, s, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	s := "1000"
	expect := true
	runSample(t, n, s, expect)
}
