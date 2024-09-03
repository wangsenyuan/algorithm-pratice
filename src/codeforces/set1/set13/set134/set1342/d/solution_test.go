package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, m []int, c []int, expect int) {
	res := solve(m, c)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	k := len(c)
	cnt := make([]int, k+1)
	check := func(arr []int) bool {
		clear(cnt)
		for _, x := range arr {
			cnt[x]++
		}
		for i := k - 1; i > 0; i-- {
			cnt[i] += cnt[i+1]
		}
		for i := 1; i <= k; i++ {
			if cnt[i] > c[i-1] {
				return false
			}
		}
		return true
	}

	var mm []int

	for _, arr := range res {
		if !check(arr) {
			t.Fatalf("Sample result %v, not correct", res)
		}
		mm = append(mm, arr...)
	}

	sort.Ints(m)
	sort.Ints(mm)

	if !reflect.DeepEqual(m, mm) {
		t.Fatalf("Sample result %v, getting non correct result %v", res, mm)
	}
}

func TestSample1(t *testing.T) {
	m := []int{1, 2, 2, 3}
	c := []int{4, 1, 1}
	expect := 3
	runSample(t, m, c, expect)
}
func TestSample2(t *testing.T) {
	m := []int{5, 8, 1, 10, 8, 7}
	c := []int{6, 6, 4, 4, 3, 2, 2, 2, 1, 1}
	expect := 2
	runSample(t, m, c, expect)
}

func TestSample3(t *testing.T) {
	m := []int{1, 1, 1, 1, 1}
	c := []int{5}
	expect := 1
	runSample(t, m, c, expect)
}

func TestSample4(t *testing.T) {
	m := []int{1, 1, 1, 1, 1}
	c := []int{1}
	expect := 5
	runSample(t, m, c, expect)
}
