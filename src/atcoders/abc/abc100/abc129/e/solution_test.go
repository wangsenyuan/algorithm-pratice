package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "10"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "1111111111111111111"
	expect := 162261460
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11"
	// (00, 00), (00, 01), (00, 10), (00, 11)
	// (01, 00), (01, 10)
	// (10, 00), (10, 01)
	// (11, 00)
	expect := 9
	// (00, 01), (01, 00)
	// (00, 11), (01, 10), (10, 01), (11, 00)
	//  少了 （10，00) 和 （00， 10）（这两个没法保证 a + b <= s)
	// 如果 a[i] = 1 (或者 b[i] = 1), 那只能是它后面的1上做文章
	// 那么有多少个个组合 (a, b) <= s[i]? 这个好像是个子问题？

	runSample(t, s, expect)
}
