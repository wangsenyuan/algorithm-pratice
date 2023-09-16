package main

import (
	"sort"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)

	if (res != "-1") != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}
	if !expect {
		return
	}
	for i := 0; i < 8; i++ {
		if codechef[i] == res[i] {
			t.Fatalf("Sample result %s, not correct", res)
		}
	}
	a := []byte(s)
	b := []byte(res)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	if string(a) != string(b) {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "codecook"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "odeckooc"
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ffffffff"
	expect := false
	runSample(t, s, expect)
}
