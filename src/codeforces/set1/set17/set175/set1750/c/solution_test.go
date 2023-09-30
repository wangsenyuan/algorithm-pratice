package main

import "testing"

func runSample(t *testing.T, a, b string, expect bool) {
	ok, res := solve(a, b)

	if expect != ok {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !expect {
		return
	}
	n := len(a)
	fa := make([]int, n+2)
	fb := make([]int, n+2)

	for _, cur := range res {
		l, r := cur[0], cur[1]
		fa[l] ^= 1
		fa[r+1] ^= 1
		fb[0] ^= 1
		fb[l] ^= 1
		fb[r+1] ^= 1
	}

	for i := 1; i <= n; i++ {
		fa[i] ^= fa[i-1]
		x := int(a[i-1]-'0') ^ fa[i]
		fb[i] ^= fb[i-1]
		y := int(b[i-1]-'0') ^ fb[i]
		if x != y || x != 0 {
			t.Fatalf("Sample result %v, not giving right answer", res)
		}
	}
}

func TestSample1(t *testing.T) {
	a := "010"
	b := "101"
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "11"
	b := "10"
	expect := false
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "1000"
	b := "0011"
	expect := false
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := "10"
	b := "10"
	expect := true
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := "111"
	b := "111"
	expect := true
	runSample(t, a, b, expect)
}
