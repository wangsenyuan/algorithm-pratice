package main

import "testing"

func runSample(t *testing.T, cities [][]int, c []int, k []int, expect int) {
	res, stations, edges := solve(cities, c, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
	n := len(cities)
	set := NewDSU(n + 1)

	var cost int
	for _, u := range stations {
		u--
		cost += c[u]
		set.Union(u, n)
	}

	for _, cur := range edges {
		i, j := cur[0], cur[1]
		i--
		j--
		dx := abs(cities[i][0] - cities[j][0])
		dy := abs(cities[i][1] - cities[j][1])
		cost += (dx + dy) * (k[i] + k[j])
		set.Union(i, j)
	}

	if cost != expect {
		t.Fatalf("Sample result %v, %v, not correct", stations, edges)
	}

	for i := 0; i < n; i++ {
		if set.Find(i) != set.Find(n) {
			t.Fatalf("Sample result %v, %v, not correct", stations, edges)
		}
	}

}

func TestSample1(t *testing.T) {
	cities := [][]int{
		{2, 3},
		{1, 1},
		{3, 2},
	}
	c := []int{3, 2, 3}
	k := []int{3, 2, 3}
	expect := 8
	runSample(t, cities, c, k, expect)
}

func TestSample2(t *testing.T) {
	cities := [][]int{
		{1, 2},
		{2, 1},
		{3, 3},
	}
	c := []int{23, 2, 23}
	k := []int{3, 2, 3}
	expect := 27
	runSample(t, cities, c, k, expect)
}
