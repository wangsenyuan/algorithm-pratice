package p786

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, K int, expect []int) {
	res := kthSmallestPrimeFraction(A, K)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %v %d, expect %v, but got %v", A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 5}
	K := 3
	expect := []int{2, 5}
	runSample(t, A, K, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 7}
	K := 1
	expect := []int{1, 7}
	runSample(t, A, K, expect)
}
