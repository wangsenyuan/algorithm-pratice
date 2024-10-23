package main

import "testing"

func runSample(t *testing.T, boxes [][]int, expect bool) {
	res := solve(boxes)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	box_sum := make([]int, len(boxes))
	var sum int
	for i, box := range boxes {
		for _, num := range box {
			box_sum[i] += num
		}
		sum += box_sum[i]
	}

	sum /= len(boxes)
	vis := make([]int, len(boxes))
	for i, cur := range res {
		if len(cur) == 0 {
			t.Fatalf("Sample result %v, not correct", res)
		}
		x, j := cur[0], cur[1]
		j--
		if i == j {
			// box_sum no change
			vis[i] += 2
			continue
		}
		ok := false
		for _, w := range cur {
			if w == x {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatalf("Sample result %v, not correct, %d gives %d, which it donest have", res, i, x)
		}
		box_sum[i] -= x
		vis[i]++
		box_sum[j] += x
		vis[j]++
	}

	for i := range len(boxes) {
		if box_sum[i] != sum {
			t.Fatalf("Sample result %v, getting %d sum %d, not expected %d", res, i, box_sum[i], sum)
		}
		if vis[i] != 2 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	boxes := [][]int{
		{1, 7, 4},
		{3, 2},
		{8, 5},
		{10},
	}
	expect := true
	runSample(t, boxes, expect)
}

func TestSample2(t *testing.T) {
	boxes := [][]int{
		{3, -2},
		{-1, 5},
	}
	expect := false
	runSample(t, boxes, expect)
}

func TestSample(t *testing.T) {
	boxes := [][]int{
		{-10, 10},
		{0, -20},
	}
	expect := true
	runSample(t, boxes, expect)
}
