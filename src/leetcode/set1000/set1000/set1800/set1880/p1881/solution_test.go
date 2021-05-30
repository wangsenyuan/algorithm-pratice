package p1881

import "testing"

func runSample(t *testing.T, nums1 []int, nums2 []int, expect int) {
	res := minimumXORSum(nums1, nums2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n1 := []int{1, 2}
	n2 := []int{2, 3}
	expect := 2
	runSample(t, n1, n2, expect)
}

func TestSample2(t *testing.T) {
	n1 := []int{1, 0, 3}
	n2 := []int{5, 3, 4}
	expect := 8
	runSample(t, n1, n2, expect)
}

func TestSample3(t *testing.T) {
	n1 := []int{0}
	n2 := []int{2877579}
	expect := 2877579
	runSample(t, n1, n2, expect)
}

func TestSample4(t *testing.T) {
	n1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	n2 := []int{8295669, 8867366, 7420350, 5745242, 4606594, 1980837, 2897226, 1597708, 3725635, 6705949, 3912749, 8305105, 7546625, 6340013}
	expect := 77947068
	runSample(t, n1, n2, expect)
}

func TestSample5(t *testing.T) {
	n1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	n2 := []int{10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000, 10000000}
	expect := 140000000
	runSample(t, n1, n2, expect)
}
