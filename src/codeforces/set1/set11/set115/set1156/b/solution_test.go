package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if expect == res {
		return
	}
	if expect == "No answer" || res == "No answer" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	// has solution
	for i := 0; i+1 < len(res); i++ {
		x := int(res[i] - 'a')
		y := int(res[i+1] - 'a')
		if x+1 == y || x == y+1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
	cnt := make([]int, 26)
	for i := range len(s) {
		cnt[int(s[i]-'a')]++
		cnt[int(res[i]-'a')]--
	}

	if slices.Max(cnt) != 0 || slices.Min(cnt) != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcd"
	expect := "cadb"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "gg"
	expect := "gg"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "codeforces"
	expect := "codeforces"
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := "abaca"
	expect := "No answer"
	runSample(t, s, expect)
}