package main

import (
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, nums []int) {
	res := solve(nums)

	if res[0] != "1" {
		t.Fatalf("Sample result %v, not correct", res)
	}

	stack := make([]string, len(nums))
	var top int
	stack[top] = res[0]
	top++
	for i := 1; i < len(nums); i++ {
		x := nums[i]

		if x == 1 {
			a := res[i-1]
			b := res[i]
			if b != a+".1" {
				t.Fatalf("must begin a new level when there is 1, but got %s %s", a, b)
			}
		} else {
			// x > 1

			if getLast(res[i]) != x {
				t.Fatalf("Sample result %s, not having the correct seq %d", res[i], x)
			}

			for top > 0 && !check(stack[top-1], res[i]) {
				top--
			}
			if top == 0 {
				t.Fatalf("Sample result %v, run out of stack", res)
			}
		}
		stack[top] = res[i]
		top++
	}

}

func check(a, b string) bool {
	return getLast(a)+1 == getLast(b)
}

func getLast(s string) int {
	i := strings.LastIndex(s, ".")
	v, _ := strconv.Atoi(s[i+1:])
	return v
}

func TestSample1(t *testing.T) {
	nums := []int{1, 1, 2, 3}
	runSample(t, nums)
}

func TestSample2(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 1, 2, 1, 2}

	runSample(t, nums)
}
