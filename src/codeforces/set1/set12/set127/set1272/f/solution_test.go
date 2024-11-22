package main

import "testing"

func runSample(t *testing.T, x, y string, expect string) {
	res := solve(x, y)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	check := func(a, b string) bool {
		for i, j := 0, 0; i < len(b); i++ {
			for j < len(a) && a[j] != b[i] {
				j++
			}
			if j == len(a) {
				return false
			}
			j++
		}
		return true
	}

	if !check(res, x) {
		t.Fatalf("Sample result %s, not correct, it don'st have %s as sub-sequence", res, x)
	}

	if !check(res, y) {
		t.Fatalf("Sample result %s, not correct, it don'st have %s as sub-sequence", res, y)
	}

	var level int

	for _, x := range []byte(res) {
		if x == '(' {
			level++
		} else {
			level--
		}
		if level < 0 {
			t.Fatalf("Sample result %s, not correct, it isn't a regular bracket sequence", res)
		}
	}
	if level != 0 {
		t.Fatalf("Sample result %s, not correct, it isn't a regular bracket sequence", res)
	}
}

func TestSample1(t *testing.T) {
	x := "(())(()"
	y := "()))()"
	expect := "(())()()"
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x := ")"
	y := "(("
	expect := "(())"
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x := ")"
	y := ")))"
	expect := "((()))"
	runSample(t, x, y, expect)
}

func TestSample4(t *testing.T) {
	x := "("
	y := "())()((((((()))(()())(())(()))))(()))()()(())((())())())(()((((())))()))))))()(()(()()(((())(()()((())()((())(((())(())((())()))(()()()()((())()())())))()()()()))())()()))(())(())()()()()()))()()))()"
	expect := "(((((((((())()((((((()))(()())(())(()))))(()))()()(())((())())())(()((((())))()))))))()(()(()()(((())(()()((())()((())(((())(())((())()))(()()()()((())()())())))()()()()))())()()))(())(())()()()()()))()()))()"
	runSample(t, x, y, expect)
}
