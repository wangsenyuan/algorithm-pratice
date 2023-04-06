package main

import (
	rand2 "math/rand"
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int) {
	n := len(P)
	R := make([]int, n)
	var cnt int
	ask := func(Q []int) []int {
		cnt++
		for i := 0; i < n; i++ {
			R[i] = P[Q[i]-1]
		}
		var X []int
		X = append(X, 1)
		mx := R[0]
		for i := 1; i < n; i++ {
			if mx < R[i] {
				X = append(X, i+1)
				mx = R[i]
			}
		}
		return X
	}

	res := solve(n, ask)

	if cnt >= n {
		t.Fatalf("Sample result took %d check", cnt)
	}

	if !reflect.DeepEqual(res, P) {
		t.Fatalf("Sample expect %v, but got %v", P, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{2, 4, 3, 1}
	runSample(t, P)
}

func TestSample2(t *testing.T) {
	n := 10

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	for x := 0; x < 5; x++ {
		rand2.Shuffle(n, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		runSample(t, arr)
	}
}

func TestSample3(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4}
	runSample(t, arr)
}

func TestSample4(t *testing.T) {
	n := 100

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	for x := 0; x < 10; x++ {
		rand2.Shuffle(n, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		runSample(t, arr)
	}
}

func TestSample5(t *testing.T) {
	n := 500

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	for x := 0; x < 30; x++ {
		rand2.Shuffle(n, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		runSample(t, arr)
	}
}
