package main

import "testing"

func runSample(t *testing.T, s string, k int, expect bool) {
	res := solve(s, k)

	if res > 0 != expect {
		t.Fatalf("Sample expect %t, but got %d", expect, res)
	}
	if !expect {
		return
	}
	n := len(s)

	buf := []byte(s)
	reverse(buf[:res])
	tmp := make([]byte, len(s))
	copy(tmp, s[res:])
	copy(tmp[n-res:], buf)

	x := string(tmp)
	for i := 0; i < n; i += k {
		for j := i; j < i+k; j++ {
			if x[j] != x[i] {
				t.Fatalf("Sample result %v, not a k-proper string", res)
			}
		}
		if i > 0 && x[i] == x[i-1] {
			t.Fatalf("Sample result %v, not a k-proper string", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "11100001"
	k := 4
	expect := true
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1110"
	k := 2
	expect := false
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "111000100011"
	k := 3
	expect := true
	runSample(t, s, k, expect)
}
func TestSample4(t *testing.T) {
	s := "00000"
	k := 5
	expect := true
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "101001"
	k := 1
	expect := true
	runSample(t, s, k, expect)
}

func TestSample6(t *testing.T) {
	s := "110001100110"
	k := 2
	expect := true
	runSample(t, s, k, expect)
}

func TestSample7(t *testing.T) {
	s := "01110001"
	k := 4
	expect := false
	runSample(t, s, k, expect)
}
