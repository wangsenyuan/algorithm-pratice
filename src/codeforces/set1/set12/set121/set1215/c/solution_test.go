package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool, correct int) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, y, ok, res := process(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}
	if len(res) != correct {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	u := []byte(x)
	v := []byte(y)
	for _, cur := range res {
		i, j := cur[0]-1, cur[1]-1
		u[i], v[j] = v[j], u[i]
	}

	x = string(u)
	y = string(v)

	if x != y {
		t.Fatalf("Sample result %v, not correct, it getting %s\n%s", res, x, y)
	}
}

func TestSample1(t *testing.T) {
	s := `4
abab
aabb
`
	expect := true
	correct := 2
	runSample(t, s, expect, correct)
}

func TestSample2(t *testing.T) {
	s := `1
a
b
`
	expect := false
	correct := -1
	runSample(t, s, expect, correct)
}

func TestSample3(t *testing.T) {
	s := `8
babbaabb
abababaa
`
	expect := true
	correct := 3
	runSample(t, s, expect, correct)
}
func TestSample4(t *testing.T) {
	s := `2
ab
ba
`
	expect := true
	correct := 2
	runSample(t, s, expect, correct)
}
