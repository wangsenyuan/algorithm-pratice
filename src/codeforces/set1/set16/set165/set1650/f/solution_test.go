package main

import "testing"

func runSample(t *testing.T, A []int, ops [][]int, expect bool) {
	res := solve(A, ops)
	if !expect {
		if len(res) != 0 {
			t.Fatalf("Sample expect no result, but got %v", res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect %t, but got none", expect)
		return
	}

	n := len(A)
	completeAt := make([]int, n)
	completePercent := make([]int, n)
	var cur int
	for _, i := range res {
		i--
		op := ops[i]
		j := op[0] - 1
		cur += op[1]
		completeAt[j] = cur
		completePercent[j] += op[2]
	}

	for i := 0; i < n; i++ {
		if completeAt[i] > A[i] || completePercent[i] < 100 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 7, 8}
	ops := [][]int{
		{1, 1, 30},
		{2, 3, 50},
		{2, 3, 100},
		{1, 1, 80},
		{3, 3, 100},
	}
	expect := true
	runSample(t, A, ops, expect)
}

func TestSample2(t *testing.T) {
	A := []int{51}
	ops := [][]int{
		{1, 36, 91},
		{1, 8, 40},
		{1, 42, 83},
		{1, 3, 45},
		{1, 13, 40},
	}
	expect := true
	runSample(t, A, ops, expect)
}

func TestSample3(t *testing.T) {
	A := []int{9, 20}
	ops := [][]int{
		{2, 8, 64},
		{2, 7, 64},
		{1, 20, 56},
		{2, 8, 76},
		{2, 20, 48},
		{1, 2, 89},
		{1, 3, 38},
		{2, 18, 66},
		{1, 7, 51},
	}
	expect := true
	runSample(t, A, ops, expect)
}

func TestSample4(t *testing.T) {
	A := []int{7, 18, 33}
	ops := [][]int{
		{1, 5, 80},
		{3, 4, 37},
	}
	expect := false
	runSample(t, A, ops, expect)
}

func TestSample5(t *testing.T) {
	A := []int{569452312, 703565975}
	ops := [][]int{
		{1, 928391659, 66},
		{1, 915310, 82},
		{2, 87017081, 92},
		{1, 415310, 54},
		{2, 567745964, 82},
	}
	expect := true
	runSample(t, A, ops, expect)
}

func TestSample6(t *testing.T) {
	A := []int{20, 31, 40}
	ops := [][]int{
		{1, 9, 64},
		{3, 17, 100},
		{3, 9, 59},
		{3, 18, 57},
		{3, 20, 49},
		{2, 20, 82},
		{2, 14, 95},
		{1, 8, 75},
		{2, 16, 67},
	}
	expect := false
	runSample(t, A, ops, expect)
}

func TestSample7(t *testing.T) {
	A := []int{20, 36}
	ops := [][]int{
		{2, 2, 66},
		{2, 20, 93},
		{1, 3, 46},
		{1, 10, 64},
		{2, 8, 49},
		{2, 18, 40},
	}
	expect := true
	runSample(t, A, ops, expect)
}

func TestSample8(t *testing.T) {
	A := []int{1000000000}
	ops := [][]int{
		{1, 1000000000, 100},
	}
	expect := true
	runSample(t, A, ops, expect)
}
