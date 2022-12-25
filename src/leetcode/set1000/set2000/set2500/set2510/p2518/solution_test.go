package p2518

import "testing"

func runSample(t *testing.T, nums []int, k int, expect int) {
	res := countPartitions(nums, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	k := 4
	expect := 6
	runSample(t, nums, k, expect)
}

func TestSample2(t *testing.T) {
	nums := []int{3, 3, 3}
	k := 4
	expect := 0
	runSample(t, nums, k, expect)
}

func TestSample3(t *testing.T) {
	nums := []int{6, 6}
	k := 2
	expect := 2
	runSample(t, nums, k, expect)
}

func TestSample4(t *testing.T) {
	nums := []int{977208288, 291246471, 396289084, 732660386, 353072667, 34663752, 815193508, 717830630, 566248717, 260280127, 824313248, 701810861, 923747990, 478854232, 781012117, 525524820, 816579805, 861362222, 854099903, 300587204, 746393859, 34127045, 823962434, 587009583, 562784266, 115917238, 763768139, 393348369, 3433689, 586722616, 736284943, 596503829, 205828197, 500187252, 86545000, 490597209, 497434538, 398468724, 267376069, 514045919, 172592777, 469713137, 294042883, 985724156, 388968179, 819754989, 271627185, 378316864, 820060916, 436058499, 385836880, 818060440, 727928431, 737435034, 888699172, 961120185, 907997012, 619204728, 804452206, 108201344, 986517084, 650443054}
	k := 95
	expect := 145586000
	runSample(t, nums, k, expect)
}
