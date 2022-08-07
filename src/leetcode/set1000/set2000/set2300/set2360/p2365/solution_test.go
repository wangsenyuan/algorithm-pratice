package p2365

import "testing"

func runSample(t *testing.T, nums []int, expect int64) {
	res := minimumReplacement(nums)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{12, 9, 7, 6, 17, 19, 21}
	expect := int64(6)
	runSample(t, nums, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{368, 112, 2, 282, 349, 127, 36, 98, 371, 79, 309, 221, 175, 262, 224, 215, 230, 250, 84, 269, 384, 328, 118, 97, 17, 105, 342, 344, 242, 160, 394, 17, 120, 335, 76, 101, 260, 244, 378, 375, 164, 190, 320, 376, 197, 398, 353, 138, 362, 38, 54, 172, 3, 300, 264, 165, 251, 24, 312, 355, 237, 314, 397, 101, 117, 268, 36, 165, 373, 269, 351, 67, 263, 332, 296, 13, 118, 294, 159, 137, 82, 288, 250, 131, 354, 261, 192, 111, 16, 139, 261, 295, 112, 121, 234, 335, 256, 303, 328, 242, 260, 346, 22, 277, 179, 223}
	expect := int64(17748)
	runSample(t, nums, expect)
}
