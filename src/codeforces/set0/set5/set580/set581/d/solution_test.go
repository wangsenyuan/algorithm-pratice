package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, pos []int, expect bool) {
	res := solve(pos)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	n := len(res)
	if n != len(res[0]) {
		t.Fatalf("Sample result %v, not a square", res)
	}

	ans := make([]int, 6)

	for i := 0; i < n; i++ {
		for j := 0; j < n; {
			k := j
			for j < n && res[i][j] == res[i][k] {
				j++
			}
			x := int(res[i][k] - 'A')
			ans[2*x+1] = j - k
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; {
			k := i
			for i < n && res[k][j] == res[i][j] {
				i++
			}
			x := int(res[k][j] - 'A')
			ans[2*x] = i - k
		}
	}

	for i := 0; i < 3; i++ {
		if ans[2*i] > ans[2*i+1] {
			ans[2*i], ans[2*i+1] = ans[2*i+1], ans[2*i]
		}
	}

	if !reflect.DeepEqual(pos, ans) {
		t.Fatalf("Sample result %v, not correct, not getting the correct position, expect %v, but got %v", res, pos, ans)
	}

}

func TestSample1(t *testing.T) {
	pos := []int{5, 1, 2, 5, 5, 2}
	expect := true
	runSample(t, pos, expect)
}

func TestSample2(t *testing.T) {
	pos := []int{4, 4, 2, 6, 4, 2}
	expect := true
	runSample(t, pos, expect)
}
