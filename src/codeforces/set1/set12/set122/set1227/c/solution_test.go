package main

import "testing"

func runSample(t *testing.T, s string, k int) {
	res := solve(s, k)

	buf := []byte(s)

	for _, cur := range res {
		l, r := cur[0], cur[1]
		reverse(buf[l-1 : r])
	}

	var level int
	var x int
	for i := 0; i < len(buf); i++ {
		if buf[i] == '(' {
			level++
		} else {
			level--
		}
		if level < 0 {
			t.Fatalf("Sample result %v, not correct, leading to a non-regular bracket string %s", res, string(buf))
		}
		if level == 0 {
			x++
		}
	}

	if x != k || level != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "()(())()"
	k := 2
	runSample(t, s, k)
}

func TestSample2(t *testing.T) {
	s := "))()()()(("
	k := 3
	runSample(t, s, k)
}
