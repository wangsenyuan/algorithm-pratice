package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if expect == res {
		return
	}

	if expect == "-" || res == "-" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	// 肯定期望一个结果
	var buf []byte
	for i, x := range []byte(s) {
		if res[i] == '1' {
			buf = append(buf, x)
		}
	}
	for i, x := range []byte(s) {
		if res[i] == '2' {
			buf = append(buf, x)
		}
	}

	for i := 1; i < len(buf); i++ {
		if buf[i-1] > buf[i] {
			t.Fatalf("Sample result %s, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := "040425524644"
	expect := "121212211211"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "123456789"
	expect := "222222222"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "98"
	expect := "21"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "987"
	expect := "-"
	runSample(t, s, expect)
}
