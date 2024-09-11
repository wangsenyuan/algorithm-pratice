package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if !expect {
		return
	}

	pos := make([]int, 26)
	for i, x := range []byte(res) {
		pos[int(x-'a')] = i
	}

	cur := pos[int(s[0]-'a')]

	for i := 1; i < len(s); i++ {
		x := int(s[i] - 'a')
		if abs(pos[x]-cur) != 1 {
			t.Fatalf("Sample result %s, not correct", res)
		}
		cur = pos[x]
	}
}

func TestSample1(t *testing.T) {
	s := "ababa"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "codedoca"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abcda"
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "zxzytyz"
	expect := true
	runSample(t, s, expect)
}
