package main

import "testing"

func runSample(t *testing.T, s string, x string, expect string) {
	res := solve(s, x)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	if expect == "-" {
		if res != "-" {
			t.Fatalf("Sample expect %s, but got %s", expect, res)
		}
		return
	}
	for i, j := 0, 0; i < len(res); i++ {
		for j < len(s) && res[i] != s[j] {
			j++
		}
		if j == len(s) {
			t.Fatalf("Sample result %s, not a sub sequence of %s", res, s)
		}
		j++
	}
	// 检查res是t的删除了连续段后的结果
	var l int
	for l < len(res) && res[l] == x[l] {
		l++
	}
	if l == len(res) {
		// 这个是前缀, ok
		return
	}
	var r int
	for r+l < len(res) && res[len(res)-1-r] == x[len(x)-1-r] {
		r++
	}

	if l+r != len(res) {
		t.Fatalf("Sample result %s, not derived from %s", res, x)
	}
}

func TestSample1(t *testing.T) {
	s := "hi"
	x := "bob"
	runSample(t, s, x, "-")
}

func TestSample2(t *testing.T) {
	s := "abca"
	x := "accepted"
	expect := "ac"
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "abacaba"
	x := "abcdcba"
	expect := "abcba"
	runSample(t, s, x, expect)
}

func TestSample4(t *testing.T) {
	s := "fdtffutxkujflswyddvhusfcook"
	x := "kavkhnhphcvckogqqqqhdmgwjdfenzizrebefsbuhzzwhzvc"
	expect := "kvc"
	runSample(t, s, x, expect)
}
