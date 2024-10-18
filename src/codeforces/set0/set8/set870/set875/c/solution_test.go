package main

import "testing"

func runSample(t *testing.T, m int, words [][]int, expect bool) {
	ok, res := solve(m, words)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

	if !ok {
		return
	}

	flag := make([]bool, m)
	for _, x := range res {
		flag[x-1] = true
	}

	get := func(x int) int {
		if !flag[x] {
			return x
		}
		return x - m
	}

	for i := 0; i+1 < len(words); i++ {
		a := words[i]
		b := words[i+1]
		lt := false
		for j := 0; j < len(a) && j < len(b); j++ {
			u := get(a[j])
			v := get(b[j])
			if u > v {
				t.Fatalf("Sample result %v, not correct", res)
			}
			if u < v {
				lt = true
				break
			}
			// u== v
		}
		// 这里还有个问题， 如果len(a) > len(b) 咋搞呢？
		if !lt && len(a) > len(b) {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	m := 3
	words := [][]int{
		{2},
		{1},
		{1, 3, 2},
		{1, 1},
	}
	expect := true
	runSample(t, m, words, expect)
}

func TestSample2(t *testing.T) {
	m := 6
	words := [][]int{
		{1, 2},
		{1, 2},
		{1, 2, 3},
		{1, 5},
		{4, 4},
		{4, 4},
	}
	expect := true
	runSample(t, m, words, expect)
}

func TestSample3(t *testing.T) {
	m := 3
	words := [][]int{
		{3, 2, 2, 1},
		{1, 1, 3},
		{2, 3, 3},
		{3, 1},
	}
	expect := false
	runSample(t, m, words, expect)
}
