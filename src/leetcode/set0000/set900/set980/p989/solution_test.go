package p989

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, K int, expect []int) {
	res := addToArrayForm(A, K)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v %d, expect %v, but got %v", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 0, 0}
	K := 34
	expect := []int{1, 2, 3, 4}
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 7, 4}
	K := 181
	expect := []int{4, 5, 5}
	runSample(t, A, K, expect)
}

func TestSample3(t *testing.T) {
	A := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	K := 1
	expect := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	runSample(t, A, K, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0}
	K := 0
	expect := []int{0}
	runSample(t, A, K, expect)
}

func TestSample5(t *testing.T) {
	A := []int{2, 0, 6, 8, 9, 3}
	K := 237
	expect := []int{2, 0, 7, 1, 3, 0}
	runSample(t, A, K, expect)
}
