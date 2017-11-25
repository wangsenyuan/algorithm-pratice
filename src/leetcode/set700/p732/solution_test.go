package p732

import "testing"

func TestSample1(t *testing.T) {
	calendar := Constructor()

	sample := []struct {
		start  int
		end    int
		expect int
	}{
		{10, 20, 1},
		{50, 60, 1},
		{10, 40, 2},
		{5, 15, 3},
		{5, 10, 3},
		{25, 55, 3},
	}

	for _, x := range sample {
		ret := calendar.Book(x.start, x.end)
		if ret != x.expect {
			t.Errorf("after book %d - %d, the answer should be %d, but get %d", x.start, x.end, x.expect, ret)
		}
	}
}

func TestSample2(t *testing.T) {
	calendar := Constructor()

	sample := []struct {
		start int
		end   int
	}{
		{47, 50}, {1, 10}, {27, 36}, {40, 47}, {20, 27}, {15, 23}, {10, 18}, {27, 36}, {17, 25},
		{8, 17}, {24, 33}, {23, 28}, {21, 27}, {47, 50}, {14, 21}, {26, 32}, {16, 21}, {2, 7}, {24, 33},
		{6, 13}, {44, 50}, {33, 39}, {30, 36}, {6, 15}, {21, 27}, {49, 50}, {38, 45}, {4, 12}, {46, 50},
		{13, 21},
	}

	expect := []int{1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7}

	for i, x := range sample {
		ret := calendar.Book(x.start, x.end)
		if ret != expect[i] {
			t.Errorf("after book %d - %d, the answer should be %d, but get %d", x.start, x.end, expect[i], ret)
		}
	}
}
