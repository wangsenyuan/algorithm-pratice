package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res == expect {
		return
	}
	if expect == bad || res == bad {
		t.Fatalf("Sample %s, expect %s, but got %s", s, expect, res)
	}
	var level int

	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			if res[i] != s[i] {
				t.Fatalf("Sample result %s, not correct, it shouldn't change existing parenthesis at %d", res, i)
			}
		} else {
			if res[i] == '?' {
				t.Fatalf("Sample result %s, not correct, it should replace ? to parenthesis at %d", res, i)
			}
		}
		if res[i] == '(' {
			level++
		} else {
			level--
		}
		if i+1 < len(s) && level == 0 {
			t.Fatalf("Sample result %s, not correct, it shouldn't have an balanced strict prefix at %d", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "(?????"
	expect := "(()())"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "(???(???(?"
	expect := bad
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "((?()??())"
	expect := "(((()))())"
	runSample(t, s, expect)
}
