package main

import (
	"testing"
)

func runSample(t *testing.T, s string, c string, expect string) {
	res := solve(s, c)

	if expect == NA {
		if res != NA {
			t.Fatalf("Sample expect %s, but got %s", expect, res)
		}
		return
	}

	if res == NA {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	if res >= c {
		t.Fatalf("Sample result %s, not correct", res)
	}

	if res == s {
		return
	}

	buf := []byte(res)

	var l int
	for l < len(s) && s[l] == res[l] {
		l++
	}

	r := l + 1
	for r < len(s) && s[r] == res[r] {
		r++
	}

	if r == len(s) {
		t.Fatalf("Sample result %s, not a swap from %s", res, s)
	}

	buf[l], buf[r] = buf[r], buf[l]

	if string(buf) != s {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "AZAMON"
	c := "APPLE"
	expect := "AMAZON"
	runSample(t, s, c, expect)
}

func TestSample2(t *testing.T) {
	s := "AZAMON"
	c := "AAAAAAAAAAALIBABA"
	expect := NA
	runSample(t, s, c, expect)
}

func TestSample3(t *testing.T) {
	s := "APPLE"
	c := "BANANA"
	expect := "APPLE"
	runSample(t, s, c, expect)
}

func TestSample4(t *testing.T) {
	s := "OS"
	c := "KSK"
	expect := "---"
	runSample(t, s, c, expect)
}
