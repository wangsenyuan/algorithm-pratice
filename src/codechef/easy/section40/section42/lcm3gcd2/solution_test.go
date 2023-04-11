package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func runSample(t *testing.T, A []int) {
	n := len(A)

	ask := func(arr []int) string {
		if len(arr) == 2 {
			i, j := arr[0], arr[1]
			g := gcd(A[i-1], A[j-1])
			return fmt.Sprintf("%d", g)
		}
		i, j, k := arr[0], arr[1], arr[2]
		l := lcm3(A[i-1], A[j-1], A[k-1])
		return fmt.Sprintf("%d", l)
	}

	res := solve(n, ask)

	for i := 0; i < n; i++ {
		cur, _ := strconv.Atoi(res[i])
		if cur != A[i] {
			t.Fatalf("Sample result %v, not correct, expect %v", res, A)
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcm3(a, b, c int) int {
	x := lcm(a, b)
	return lcm(x, c)
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 4, 3}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	n := 10
	x := 100

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(x) + 1
	}

	runSample(t, arr)
}
