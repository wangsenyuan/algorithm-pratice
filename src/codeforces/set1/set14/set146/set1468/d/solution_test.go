package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, s []int, expect int) {
	res := solve(n, a, b, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b := 7, 3, 6
	s := []int{1, 4}
	expect := 2
	runSample(t, n, a, b, s, expect)
}

func TestSample2(t *testing.T) {
	n, a, b := 8, 6, 3
	s := []int{1, 5}
	expect := 1
	runSample(t, n, a, b, s, expect)
}

func TestSample3(t *testing.T) {
	n, a, b := 7, 3, 6
	s := []int{4, 4}
	expect := 1
	runSample(t, n, a, b, s, expect)
}

func TestSample4(t *testing.T) {
	n, a, b := 2503, 2218, 1466
	s := []int{8768, 2384, 6078, 7897, 2912, 9040, 5852, 6190, 2863, 2423, 3340, 8918, 7911, 8936, 6243, 5502, 378, 6094, 5852, 4842, 1481, 2100, 6835, 1056, 3038, 939, 778, 395, 5572, 151, 9128, 3689, 4793, 546, 964, 6606, 5907, 5688, 1583, 7759, 7080, 8634, 136, 104, 564, 7947, 1157, 5111, 8503, 4950, 5872, 7502, 5956, 2837, 776, 1893, 804, 9003, 9260, 6565, 5153, 3324, 7284, 7806, 4411, 6800, 3352, 1746, 278, 6740, 4470, 2680, 9226, 5528, 1328, 5034, 8242, 3921, 3884, 8887, 1086, 4825, 1922, 8977, 3553, 5386, 7506, 6528, 3155, 9575, 7963, 286, 9664}
	expect := 14
	runSample(t, n, a, b, s, expect)
}
