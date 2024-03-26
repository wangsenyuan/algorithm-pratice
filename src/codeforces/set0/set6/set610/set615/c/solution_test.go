package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string, x string, expect int) {
	res := solve(s, x)

	if expect < 0 && len(res) > 0 {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	if expect < 0 {
		return
	}
	if expect != len(res) {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	var buf bytes.Buffer

	for _, tmp := range res {
		l, r := tmp[0], tmp[1]
		if l > r {
			l, r = r, l
			buf.WriteString(reverse(s[l-1 : r]))
		} else {
			buf.WriteString(s[l-1 : r])
		}
	}

	y := buf.String()

	if y != x {
		t.Fatalf("Sample result %v => %s, not expect %s", res, y, x)
	}

}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func TestSample1(t *testing.T) {
	s := "abc"
	x := "cbaabc"
	expect := 2
	runSample(t, s, x, expect)
}

func TestSample2(t *testing.T) {
	s := "aaabrytaaa"
	x := "ayrat"
	expect := 3
	runSample(t, s, x, expect)
}

func TestSample3(t *testing.T) {
	s := "ami"
	x := "no"
	expect := -1
	runSample(t, s, x, expect)
}
