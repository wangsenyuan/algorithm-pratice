package main

import (
	"fmt"
	"testing"
)

func runSample(t *testing.T, expect string) {
	ask := func(arr []int) string {
		if len(arr) == 1 {
			i := arr[0] - 1
			return expect[i : i+1]
		}
		l, r := arr[0], arr[1]
		l--
		set := make(map[byte]int)
		for i := l; i < r; i++ {
			set[expect[i]]++
		}
		return fmt.Sprintf("%d", len(set))
	}

	res := solve(len(expect), ask)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	expect := "guess"
	runSample(t, expect)
}
