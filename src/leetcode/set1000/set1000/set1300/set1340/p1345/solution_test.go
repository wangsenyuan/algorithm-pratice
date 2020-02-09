package p1345

import "testing"

func runSample(t *testing.T, arr []int, expect int) {
	res := minJumps(arr)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{7}
	expect := 0
	runSample(t, arr, expect)
}

func TestSample3(t *testing.T) {
	arr := []int{7, 6, 9, 6, 9, 6, 9, 7}
	expect := 1
	runSample(t, arr, expect)
}

func TestSample4(t *testing.T) {
	arr := []int{6, 1, 9}
	expect := 2
	runSample(t, arr, expect)
}

func TestSample5(t *testing.T) {
	arr := []int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13}
	expect := 3
	runSample(t, arr, expect)
}

func TestSample6(t *testing.T) {
	arr := []int{68, -94, -44, -18, -1, 18, -87, 29, -6, -87, -27, 37, -57, 7, 18, 68, -59, 29, 7, 53, -27, -59, 18, -1, 18, -18, -59, -1, -18, -84, -20, 7, 7, -87, -18, -84, -20, -27}
	expect := 5
	runSample(t, arr, expect)
}
