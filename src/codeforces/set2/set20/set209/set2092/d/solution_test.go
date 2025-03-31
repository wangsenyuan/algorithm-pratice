package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	ok, ans := solve(s)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, ans)
	}
	if !expect {
		return
	}
	for _, i := range ans {
		i--
		if s[i] == s[i+1] {
			t.Fatalf("Sample result %v, not correct", ans)
		}
		c := base[0]
		if c == s[i] || c == s[i+1] {
			c = base[1]
		}
		if c == s[i] || c == s[i+1] {
			c = base[2]
		}
		s = s[:i+1] + string(c) + s[i+1:]
	}
	cnt := count(s, base[0])
	if count(s, base[1]) != cnt || count(s, base[2]) != cnt {
		t.Fatalf("Sample result %v, get wrong result %s", ans, s)
	}
}

func TestSmaple1(t *testing.T) {
	s := "TILII"
	expect := true
	runSample(t, s, expect)
}

func TestSmaple2(t *testing.T) {
	s := "L"
	expect := false
	runSample(t, s, expect)
}


func TestSmaple3(t *testing.T) {
	s := "LIT"
	expect := true
	runSample(t, s, expect)
}
