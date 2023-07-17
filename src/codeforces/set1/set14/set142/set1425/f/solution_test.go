package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int) {
	var cnt int
	ask := func(l int, r int) int {
		cnt++
		var sum int
		for i := l - 1; i < r; i++ {
			sum += A[i]
		}
		return sum
	}

	res := solve(len(A), ask)

	if cnt > len(A) {
		t.Fatalf("Sample result took %d times to ask > %d", cnt, len(A))
	}

	if !reflect.DeepEqual(res, A) {
		t.Fatalf("Sample expect %v, but got %v", A, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 4, 6, 7, 8}
	runSample(t, A)
}

func TestSample2(t *testing.T) {
	A := rand.Perm(100)
	runSample(t, A)
}
