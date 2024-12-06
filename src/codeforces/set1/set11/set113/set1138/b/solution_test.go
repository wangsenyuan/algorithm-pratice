package main

import "testing"

func runSample(t *testing.T, c string, a string, expect bool) {
	res := solve(c, a)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	n := len(c)

	if len(res) != n/2 {
		t.Fatalf("Sample result %v, not correct", res)
	}

	marked := make([]bool, n)
	for _, i := range res {
		marked[i-1] = true
	}
	cnt := make([]int, 2)
	for i := 0; i < n; i++ {
		if marked[i] {
			if c[i] == '1' {
				cnt[0]++
			}
		} else {
			if a[i] == '1' {
				cnt[1]++
			}
		}
	}

	if cnt[0] != cnt[1] {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	c := "0011"
	a := "0101"
	expect := true
	runSample(t, c, a, expect)
}

func TestSample2(t *testing.T) {
	c := "000000"
	a := "111111"
	expect := false
	runSample(t, c, a, expect)
}

func TestSample3(t *testing.T) {
	c := "0011"
	a := "1100"
	expect := true
	runSample(t, c, a, expect)
}


func TestSample4(t *testing.T) {
	c := "00100101"
	a := "01111100"
	expect := true
	runSample(t, c, a, expect)
}