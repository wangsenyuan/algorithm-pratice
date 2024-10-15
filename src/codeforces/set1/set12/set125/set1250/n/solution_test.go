package main

import "testing"

func runSample(t *testing.T, wires [][]int, expect int) {
	res := solve(wires)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	use := make([]int, len(wires))
	for i := 0; i < len(wires); i++ {
		use[i] = -1
	}

	for i, cur := range res {
		e_id := cur[0]
		use[e_id-1] = i
	}

	pos, arr := compact(wires)

	set := NewDSU(len(arr))

	for i, wire := range wires {
		u, v := wire[0], wire[1]
		if use[i] < 0 {
			u = pos[u]
			v = pos[v]
		} else {
			cur := res[use[i]]
			a, b := cur[1], cur[2]
			if a == u {
				u = pos[b]
				v = pos[v]
			} else {
				// a == v
				u = pos[u]
				v = pos[b]
			}
		}
		set.Union(u, v)
	}
	fa := -1
	for i := 0; i < len(arr); i++ {
		j := set.Find(i)
		if set.cnt[j] == 1 {
			continue
		}
		if j == fa {
			// good
			continue
		}
		if fa < 0 {
			fa = j
		} else {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	wires := [][]int{{4, 7}}
	expect := 0
	runSample(t, wires, expect)
}

func TestSample2(t *testing.T) {
	wires := [][]int{
		{1, 2},
		{2, 3},
		{4, 5},
		{5, 6},
	}
	expect := 1
	runSample(t, wires, expect)
}

func TestSample3(t *testing.T) {
	wires := [][]int{
		{5, 8},
		{7, 10},
		{2, 10},
		{13, 10},
		{7, 13},
	}
	expect := 1
	runSample(t, wires, expect)
}
