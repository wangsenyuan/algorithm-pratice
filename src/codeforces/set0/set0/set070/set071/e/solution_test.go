package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, source string, target string, expect bool) {
	res := solve(source, target)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	ss := strings.Split(source, " ")
	// tt := strings.Split(target, " ")
	cnt := make(map[string]int)
	for _, x := range ss {
		cnt[x]++
	}

	for it, cur := range res {
		m := len(cur)
		if m <= 1 {
			t.Fatalf("Sample result %v, not correct at %d", res, it)
		}
		var tmp int
		for i := 0; i < m-1; i++ {
			tmp += atoms[cur[i]]
			cnt[cur[i]]--
		}

		if tmp != atoms[cur[m-1]] {
			t.Fatalf("Sample result %v, not getting the correct formula at %d", res, it)
		}
	}

	for _, v := range cnt {
		if v != 0 {
			t.Fatalf("Sample result %v, not correct, having some thing wrong %v", res, cnt)
		}
	}
}

func TestSample1(t *testing.T) {
	source := "Mn Co Li Mg C P F Zn Sc K"
	target := "Sn Pt Y"
	expect := true
	runSample(t, source, target, expect)
}

func TestSample2(t *testing.T) {
	source := "H H"
	target := "He"
	expect := true
	runSample(t, source, target, expect)
}

func TestSample3(t *testing.T) {
	source := "Bk Fm"
	target := "Cf Es"
	expect := false
	runSample(t, source, target, expect)
}

func TestSample4(t *testing.T) {
	source := "Be Li He B"
	target := "N N"
	expect := true
	runSample(t, source, target, expect)
}
