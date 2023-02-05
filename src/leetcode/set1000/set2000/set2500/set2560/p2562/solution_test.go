package p2562

import "testing"

func runSample(t *testing.T, A, B []int, expect int64) {
	res := minCost(A, B)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 2, 2}
	B := []int{1, 4, 1, 2}
	var expect int64 = 1
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 4, 1}
	B := []int{3, 2, 5, 1}
	var expect int64 = -1
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88}
	B := []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8}
	var expect int64 = 48
	runSample(t, A, B, expect)
}

func TestSample4(t *testing.T) {
	A := []int{4, 4, 4, 4, 3}
	B := []int{5, 5, 5, 5, 3}
	var expect int64 = 8
	runSample(t, A, B, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1}
	B := []int{1}
	var expect int64
	runSample(t, A, B, expect)
}

func TestSample6(t *testing.T) {
	A := []int{3350, 1104, 2004, 1577, 1365, 2088, 2249, 1948, 2621, 750, 31, 2004, 1749, 3365, 3350, 3843, 3365, 1656, 3168, 3106, 2820, 3557, 1095, 2446, 573, 2464, 2172, 1326, 2712, 467, 1104, 1446, 1577, 53, 2492, 2638, 1200, 2997, 3454, 2492, 1926, 1452, 2712, 446, 2997, 2820, 750, 2529, 3847, 656, 272, 3873, 530, 1749, 1743, 251, 3847, 31, 251, 515, 2858, 126, 2491}
	B := []int{530, 1920, 2529, 2317, 1969, 2317, 1095, 2249, 2858, 2636, 3772, 53, 3106, 2638, 1267, 1926, 2882, 515, 3772, 1969, 3454, 2446, 656, 2621, 1365, 1743, 3557, 1656, 3447, 446, 1098, 1446, 467, 2636, 1088, 1098, 2882, 1088, 1326, 644, 3873, 3843, 3926, 1920, 2464, 2088, 205, 1200, 1267, 272, 925, 925, 2172, 2491, 3168, 644, 1452, 573, 1948, 3926, 205, 126, 3447}
	var expect int64 = 837
	runSample(t, A, B, expect)
}
