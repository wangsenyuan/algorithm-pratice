package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, pairs [][]int, expect []int) {

	res := solve(pairs)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

}

func abs(num int) int {
	return max(num, -num)
}

func TestSample1(t *testing.T) {
	pairs := [][]int{
		{1, 10},
		{8, 13},
		{3, 4},
		{5, 20},
	}
	expect := []int{8, 8, 4, 5}
	runSample(t, pairs, expect)
}

func TestSample2(t *testing.T) {
	pairs := [][]int{
		{335279264, 849598327},
		{446755913, 822889311},
		{526239859, 548830120},
		{181424399, 715477619},
		{342858071, 625711486},
		{448565595, 480845266},
		{467825612, 647639160},
		{160714711, 449656269},
		{336869678, 545923679},
		{61020590, 573085537},
		{626006012, 816372580},
		{135599877, 389312924},
		{511429216, 547865075},
		{561330066, 605997004},
		{539239436, 921749002},
	}
	expect := []int{526239859, 526239859, 526239859, 467825612, 467825612, 467825612, 467825612, 449656269, 449656269, 449656269, 626006012, 389312924, 511429216, 561330066, 561330066}
	runSample(t, pairs, expect)
}
